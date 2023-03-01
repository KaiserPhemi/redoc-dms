package models

import (
	"time"

	"github.com/KaiserPhemi/redoc-dms/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

// User model
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// connects to db
func init() {
	db = config.Connect()
	db.AutoMigrate(&User{})
}

// fetches all users
func FetchAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

// fetch a user by Id
func FetchUserById(Id uint64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID", Id).Limit(1).Find(&getUser)
	return &getUser, db
}

// deletes a user by Id
func DeleteUser(Id uint64) User {
	var user User
	db.Where("ID", Id).Delete(user)
	return user
}

// adds a new user
func (newUser *User) AddUser() *User {
	db.Create(&newUser)
	return newUser
}
