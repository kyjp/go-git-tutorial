package infra

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"log"
	"os"
)

func SetupDB() *gorm.DB {
	env := os.Getenv("ENV")
	println(env)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		"database",
		"postgres",
		"postgres",
		"app",
		"5432",
	)

	var (
		db *gorm.DB
		err error
	)

	if env == "test" {
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		log.Println("Setup sqlite database")
	} else {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		log.Println("Setup postgresql database")
	}
	
	if err != nil {
		panic("Failed to connect database")
	}

	return db
}