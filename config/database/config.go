package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
)

var (
	db *gorm.DB
)

func ConfigDB() {
	conn, err := DbSetup()
	if err != nil {
		panic(err)
	}
	db = conn
	fmt.Println("=======Database Connection Success=======")
}

func DatabaseMigration() {
	conn, err := DbSetup()
	if err != nil {
		panic(err)
	}
	db = conn
	DbMigration()
	fmt.Println("=======MIGRATION Table Success=======")
}

func DbMigrateFresh() {
	DropMigration()
	DatabaseMigration()

}

func DropMigration() {
	conn, err := DbSetup()
	if err != nil {
		panic(err)
	}
	db = conn
	Drop()
	fmt.Println("=======DROP Table Success=======")
}

func DbSetup() (*gorm.DB, error) {
	// MYSQL
	// dsn := "root:@tcp(127.0.0.1:3306)/github.com/fendilaw41/absensifendiray?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	QueryFields: true,
	// })
	// POSTGRESS 56324
	errenv := godotenv.Load(".env")
	if errenv != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
