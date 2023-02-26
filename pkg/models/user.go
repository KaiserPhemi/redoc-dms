package models

import (
	"time"

	"github.com/KaiserPhemi/redoc-dms/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

// User model
type User struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `json:"name"`
	Email     string `json:"email"`
	Password  string  `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// connects to db
func init() {
	db = config.Connect()
	db.AutoMigrate(&User{})
}
