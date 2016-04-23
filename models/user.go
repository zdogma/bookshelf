package models

import (
	"github.com/go-xorm/xorm"

	// 依存関係の解決
	_ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

// init ...
func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root@/bookshelf")
	if err != nil {
		panic(err)
	}
}

// User is
type User struct {
	ID       int    `json:"id" xorm:"'id'"`
	Username string `json:"name" xorm:"'name'"`
}

// NewUser ...
func NewUser(id int, name string) User {
	return User{
		ID:       id,
		Username: name,
	}
}

// UserRepository is
type UserRepository struct {
}

// NewUserRepository ...
func NewUserRepository() UserRepository {
	return UserRepository{}
}

// FindByID ...
func (model UserRepository) FindByID(id int) *User {
	var user = User{ID: id}
	has, _ := engine.Get(&user)
	if has {
		return &user
	}

	return nil
}
