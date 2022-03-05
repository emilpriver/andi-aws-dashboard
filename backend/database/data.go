package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GithubUser struct {
	gorm.Model
	ID       int
	GithubID int
	Username string
}

type User struct {
	gorm.Model
	Email    string
	Username string
	Avatar   string
	Github   GithubUser
	GithubID int
}

func Init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Stockholm",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Print("Error on auto migrate")
		return
	}

	err = db.AutoMigrate(&GithubUser{})
	if err != nil {
		fmt.Print("Error on auto migrate")
		return
	}
}
