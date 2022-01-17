package db_driver

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

func GetConnection() *gorm.DB {
	//if run in localhost
	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//	return nil
	//}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOSTNAME")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSL := os.Getenv("DB_SSL")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPass, dbPort, dbName, dbSSL)

	db, err2 := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err2 != nil {
		panic("Failed to create a connection to database")
	} else {
		dbConfig, _ := db.DB()
		dbConfig.SetMaxIdleConns(5)
		dbConfig.SetMaxOpenConns(20)
		dbConfig.SetConnMaxIdleTime(3 * time.Minute)
		dbConfig.SetConnMaxLifetime(5 * time.Minute)

	}

	fmt.Println("open connection to db")

	return db
}

func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	fmt.Println("close connection from db")
	dbSQL.Close()
}
