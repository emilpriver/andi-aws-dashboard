package github

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Andi-App/Andi/database"
	"github.com/Andi-App/Andi/integration"
	"github.com/Andi-App/Andi/utils"
	"github.com/gin-gonic/gin"
)

func GithubLoginCallback(c *gin.Context) {
	callbackCode, callbackCodeExists := c.GetQuery("code")

	if !callbackCodeExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing callback code",
		})
		return
	}

	accessToken, err := integration.
		GithubGetUserAccessToken(callbackCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Coudn't get access token from Github",
		})

		return
	}

	ghUser, err := integration.
		GithubGetUserByToken(accessToken)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error loading user.",
		})

		return
	}

	user := database.User{}

	database.DB.
		Preload("github_users").
		Table("users").
		Joins("INNER JOIN github_users ON (github_users.id = users.github_id)").
		Where("github_users.github_id = ?", ghUser.ID).
		Find(&user)

	fmt.Println(user.ID)

	if user.ID == 0 {
		newUser := &database.User{
			Email:    ghUser.Email,
			Username: ghUser.Login,
			Avatar:   ghUser.AvatarUrl,
			Github: database.GithubUser{
				GithubID: ghUser.ID,
				Username: ghUser.Login,
			},
		}

		result := database.DB.
			Create(&newUser)

		if result.Error != nil {
			panic("Couldn't create Github User")
		}

		user = *newUser
	}

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error fetching user",
		})

		return
	}

	jwt, err := utils.GenerateJWT(&user, "github")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error generating JWT token",
		})
	}

	domainSecure, _ := strconv.ParseBool(os.Getenv("FRONTEND_COOKIE_SECURE"))
	c.SetCookie(
		"token",
		jwt.Token,
		int(jwt.Expires.Sub(time.Now()).Seconds()),
		"/",
		os.Getenv("FRONTEND_DOMAIN"),
		domainSecure,
		true,
	)

	redirectURL := fmt.Sprintf(
		"%v/dashboard",
		os.Getenv("FRONTEND_DOMAIN"),
	)

	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
