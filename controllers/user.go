package controllers

import "bookshelf/models"

// User is ...
type User struct {
}

// NewUser ...
func NewUser() User {
	return User{}
}

// Find ...
func (controller User) Find(id int) interface{} {
	repo := models.NewUserRepository()
	user := repo.FindByID(id)

	return user
}
