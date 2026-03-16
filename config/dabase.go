package config

import (
	"fmt"
	"os"

	"studentapi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	user     := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host     := os.Getenv("DB_HOST")
	port     := os.Getenv("DB_PORT")
	dbname   := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		user, password, host, port, dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ Database connect nahi hua: " + err.Error())
	}

	// Students table automatically banega
	db.AutoMigrate(&models.Student{})

	DB = db
	fmt.Println("✅ Database connected!")
}

/*
**Aur `.env` file banao:**
DB_USER=root
DB_PASSWORD=tumhara_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=studentdb
*/