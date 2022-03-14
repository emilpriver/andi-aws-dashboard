package database

import "gorm.io/gorm"

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
	Teams    []*Team `gorm:"many2many:user_teams;"`
}
