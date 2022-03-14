package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GithubUser struct {
	gorm.Model
	GithubID int64 `gorm:"uniqueIndex"`
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

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Stockholm",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	initDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = initDB

	err = DB.AutoMigrate(&User{})
	if err != nil {
		fmt.Print("Error on auto migrate")
		return
	}

	err = DB.AutoMigrate(&GithubUser{})
	if err != nil {
		fmt.Print("Error on auto migrate")
		return
	}
}
