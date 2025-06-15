package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// Загружаем файл .env
	err := godotenv.Load()
	if err != nil {
		panic("Ошибка загрузки .env файла: " + err.Error())
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DATABASE |Failed to connect to database: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("DATABASE |Error getting base: " + err.Error())
	}
	if err := sqlDB.Ping(); err != nil {
		panic("DATABASE |Failed to connect to database: " + err.Error())
	}

	fmt.Println("DATABASE | Successful connection to the database!")
	return db
}
