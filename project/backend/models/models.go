package models

import(
		     //grom
		 "github.com/jinzhu/gorm"
		 _ "github.com/jinzhu/gorm/dialects/postgres"   
	)

	type Company struct {
		gorm.Model
		Name  string `gorm:"unique;not null"`
		Admin Admin
	}
	
	type Admin struct {
		gorm.Model
		Username  string  `gorm:"unique;not null"`
		Email     string  `gorm:"unique;not null"`
		Password  []byte  `gorm:"not null"`
		CompanyID uint    `gorm:"index"`
	}
	
	type User struct {
		gorm.Model
		Username  string  `gorm:"unique;not null"`
		Email     string  `gorm:"unique;not null"`
		Password  []byte  `gorm:"not null"`
		CompanyID uint    `gorm:"index"`
	}
	
