package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/Andi-App/Andi/database"
	"github.com/dgrijalva/jwt-go"
)

type Jwt struct {
	Token   string
	Expires time.Time
}

func GenerateJWT(user *database.User, method string) (*Jwt, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	var mySigningKey = []byte(jwtSecret)
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)

	exp := time.Now().Add(time.Hour * 12)

	claims["authorized"] = true
	claims["userID"] = user.ID
	claims["authMethod"] = method
	claims["exp"] = exp.Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Error generating JWT token: %s", err.Error())

		return &Jwt{}, err
	}

	jwtToken := Jwt{
		Token:   tokenString,
		Expires: exp,
	}

	return &jwtToken, nil
}
