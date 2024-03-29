package config

// imports
import (
	util "github.com/KaiserPhemi/redoc-dms/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// creates a database connection
func Connect() *gorm.DB {
	dbUrl := "host=localhost user=oluwafemiakinwa dbname=redoc_dev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	util.HandleError(err)
	return db
}
