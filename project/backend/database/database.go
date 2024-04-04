package database

import (
	//grom
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/fmilheir/final_year_project/backend/models"
	//"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var  DB Dbinstance

func Connect() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Paris",os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), "mydb",)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connection Opened to Database")
	db.AutoMigrate(&models.Company{})
	db.AutoMigrate(&models.Admin{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Incident{}, &models.EntityRef{}, &models.CharacteristicRelationship{}, &models.ResourceEntity{}, &models.Characteristic{})
	log.Println("Database Migrated")
	DB = Dbinstance{Db: db} 
}