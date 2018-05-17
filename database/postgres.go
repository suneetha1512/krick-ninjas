package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/xid"
	"github.com/suneetha1512/krick-ninjas/models"
)

type DB struct {
	db *gorm.DB
}

func Setup(connString string) (DB, error) {
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		return DB{}, fmt.Errorf("unable to initaizile DB %s", err)
	}
	// create tables
	db.AutoMigrate(&models.Person{}, &models.Group{})
	return DB{db}, nil;
}

func (db DB) AddPerson(person models.Person) {
	person.Id = xid.New().String()
	db.db.Create(&person)
}

func (db DB) FindPersons(person models.Person) []models.Person {
	var results []models.Person
	db.db.Where(&models.Person{FirstName: person.FirstName, LastName: person.LastName}).Find(&results)
	return results
}
