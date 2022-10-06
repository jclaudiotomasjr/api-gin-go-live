package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jclaudiotomasjr/api-gin-go-live/api-go-gin-go-live-quiz/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConectaBD() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Println("Cannot conect to database", Dbdriver)
		log.Fatal("Connection error: ", err)
	} else {
		fmt.Println("We are connected to the database: ", Dbdriver)
	}

	DB.AutoMigrate(&models.Score{})

	/*DB, err := gorm.Open(sqlite.Open("data/scores.db"), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com Banco de dados!", err)
	}
	*/
}
