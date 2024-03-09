package models

import(
	 "math/rand"
	     //grom
		 "github.com/jinzhu/gorm"
		 _ "github.com/jinzhu/gorm/dialects/postgres"   
	)



type Company struct {
    gorm.Model
	ID    int    `db:"id" gorm:"primary_key"`
	Name  string `db:"name"`
	admin string `db:"admin" gorm:unique;`
}

type Users struct {
    gorm.Model
	ID          int    `JSON:"ID" db:"id" gorm:"primary_key"`
	Username    string `JSON:"Username" gorm: unique`
	Email       string `JSON:"Email" gorm: unique`
	Password    []byte `JSON:"-" db:"password"`
	persmission string `db:"permission"`
	CompanyID   int    `JSON:"CompanyID" db:"company_id" gorm:"foreignkey:ID;association_foreignkey:ID"`
}

type Admin struct {
    gorm.Model
    ID       int    `JSON:"ID" gorm:"primary_key"`
    Username string `JSON:"Username" gorm: unique`
    Email    string `JSON:"Email" gorm: unique`
    Password string `JSON:"-" db:"password"`
    CompanyID int   `JSON:"CompanyID" db:"company_id" gorm:"foreignkey:ID;association_foreignkey:ID"`
}

func NewCompany(name, admin string) *Company {
	return &Company{
		ID: rand.Intn(1000),
		Name: name,
		admin: admin,
	}
}

func NewUser(username, email, password, permission string, companyID int) *Users {
	return &Users{
		ID: rand.Intn(1000),
		Username: username,
		Email: email,
        //cast the password to a byte array
		Password: []byte(password),
		persmission: permission,
		CompanyID: companyID,
	}
}