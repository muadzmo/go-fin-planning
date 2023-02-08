package database

import (
	"os"

	"github.com/muadzmo/go-fin-planning/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db_uname := os.Getenv("db_uname")
	db_pass := os.Getenv("db_pass")
	db_host := os.Getenv("db_host")
	db_string := db_uname + ":" + db_pass + "@tcp(" + db_host + ")"

	connection, err := gorm.Open(mysql.Open(db_string+"/fin_planning?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.Balance{}, &models.Planning{}, &models.Transaction{})
}
