package database

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fendilaw41/absensifendiray/config/database/seeds"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
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

func DbMigrateFreshSeed() {
	DropMigration()
	DatabaseMigration()
	DatabaseSeeder()
}

func DatabaseSeeder() {
	flag.Parse()
	args := flag.Args()
	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			db, err := DbSetup()
			if err != nil {
				log.Fatalf("Error opening DB: %v", err)
			}
			seeds.Execute(db, args[1:]...)
			os.Exit(0)
			fmt.Println("=======Seeders Success=======")
		}
	}
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
	// dbHost := os.Getenv("DB_HOST")
	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// dbPort := os.Getenv("DB_PORT")

	dbHost := os.Getenv("DB_HOST_SERVER")
	dbUser := os.Getenv("DB_USER_SERVER")
	dbPass := os.Getenv("DB_PASSWORD_SERVER")
	dbName := os.Getenv("DB_NAME_SERVER")
	dbPort := os.Getenv("DB_PORT_SERVER")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
