package userService

import (
	"gorm.io/gorm"
)

type UserRepoistory interface {
	GetUsers() (User, error)
	PostUser(user User) error
	PatchUserByID(userID int, user User) error
	DeleteUserByID(userID int) error
}

type userRepoistory struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepoistory {
	return &userRepoistory{db}
}

func (r *userRepoistory) GetUsers() ([]User, error) {
	var users []User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *userRepoistory) PostUser(user User) error {
	result := r.db.Create(&user)
	return result.Error
}

func (r *userRepoistory) PatchUserByID(userID int, user User) error {
	result := r.db.Model(&User{}).Where("id = ?", userID).Updates(user)
	return result.Error
}

func (r *userRepoistory) DeleteUserByID(userID int) error {
	result := r.db.Delete(&User{}, userID)
	return result.Error
}
