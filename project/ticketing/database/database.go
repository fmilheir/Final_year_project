package database

import (
	//grom
	"fmt"
	"log"
	"os"
	

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmilheir/test_swagger/models"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var  DB Dbinstance

func Connect() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Paris",os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), "mydb",)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connection Opened to Database")
	db.AutoMigrate(&models.Incident{}, &models.ResourceEntity{})
	log.Println("Database Migrated")
	DB = Dbinstance{Db: db} 
}