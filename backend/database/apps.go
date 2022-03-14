package database

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Users []*User `gorm:"many2many:user_teams;"`
	Slug  string  `gorm:"uniqueIndex"`
	Name  string
	Apps  []*App `gorm:"many2many:teams_apps;"`
}

type App struct {
	gorm.Model
	Name  string
	Slug  string  `gorm:"uniqueIndex"`
	Teams []*Team `gorm:"many2many:teams_apps;"`
}
