package github

import (
	"errors"
	"net/http"

	"github.com/Andi-App/Andi/database"
	"github.com/Andi-App/Andi/integration"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	user := database.GithubUser{}

	result := database.DB.First(&user, "github_id = ?", ghUser.ID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		newUser := &database.User{
			Email:    ghUser.Email,
			Username: ghUser.Login,
			Avatar:   ghUser.AvatarUrl,
			Github: database.GithubUser{
				GithubID: ghUser.ID,
				Username: ghUser.Login,
			},
		}

		result := database.DB.Create(&newUser)

		if result.Error != nil {
			panic("Couldn't create Github User")
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"UserID": user.Username,
	})
}
