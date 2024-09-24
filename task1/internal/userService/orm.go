package userservice

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int
	Email    string
	Password string
}
