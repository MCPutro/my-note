package test

import (
	"fmt"
	"github.com/MCPutro/my-note/entity"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"testing"
	"time"
)

func TestConnection(t *testing.T) {
	godotenv.Load("D:\\Belajar\\Go\\src\\my_note\\.env")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOSTNAME")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSL := os.Getenv("DB_SSL")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPass, dbPort, dbName, dbSSL)

	t.Run("connection db", func(t *testing.T) {

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		assert.NotNil(t, db)
		assert.Nil(t, err)

		err = db.AutoMigrate(entity.Note{})
		assert.Nil(t, err)
	})

	t.Run("test insert", func(t *testing.T) {
		db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		user := entity.User{
			Email:    "test@email.com",
			Password: "password123",
		}

		create := db.Create(&user)

		assert.Nil(t, create.Error)

		//db.Migrator().DropTable(&entity.Note{})
		//db.Migrator().DropTable(&entity.User{})

	})

	//t.Run("drop table", func(t *testing.T) {
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.Migrator().DropTable(&entity.Note{})
	db.Migrator().DropTable(&entity.User{})
	//
	//})

}

func TestCreateNore(t *testing.T) {
	godotenv.Load("D:\\Belajar\\Go\\src\\my_note\\.env")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOSTNAME")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSL := os.Getenv("DB_SSL")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPass, dbPort, dbName, dbSSL)

	t.Run("insert note", func(t *testing.T) {

		db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		newNote := entity.Note{
			Text:        "test 1111",
			Visible:     true,
			UserId:      "36929f54-6554-46d4-9e29-67ac129077f2",
			CreatedDate: time.Now(),
			UpdateDate:  time.Now(),
		}

		db.Create(&newNote)

		fmt.Println(newNote)

	})
}
