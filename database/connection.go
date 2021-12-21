package connection

import (
	"Chatapp/models"
	"Chatapp/utils"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDataBase(config *utils.Config) error {
	db, err := gorm.Open("sqlite3", config.DBName)

	if err != nil {
		log.Fatal(err)
		return err
	}

	DB = db
	return nil
}

func RunMigration(config *utils.Config) error {
	err := ConnectDataBase(config)

	if err != nil {
		log.Fatal(err)
		return err
	}

	DB.AutoMigrate(&models.User{}, &models.Chat{}, &models.Message{})
	fmt.Println("Database tables created successfully")

	return nil
}
