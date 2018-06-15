package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/xid"
	"github.com/suneetha1512/let-me-choose/models"
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
	db.AutoMigrate(&models.Person{}, &models.Choice{})
	return DB{db}, nil;
}

func (db DB) AddPerson(person models.Person) {
	person.Id = xid.New().String()
	db.db.Create(&person)
}

func (db DB) UpdatePersonChoice(memberId string, choiceId string) {
	person := db.FindPersonById(memberId)
	person.Choices = append(person.Choices, choiceId)
	db.db.Model(&person).Updates(models.Person{Choices: person.Choices})
}

func (db DB) FindPersons(person models.Person) []models.Person {
	var results []models.Person
	db.db.Where(&models.Person{FirstName: person.FirstName, LastName: person.LastName}).Find(&results)
	return results
}

func (db DB) FindPersonById(memberId string) models.Person {
	var person models.Person
	db.db.Where(&models.Person{Id: memberId}).Find(&person)
	return person
}

func (db DB) AddChoice(choice models.Choice) {
	choice.Id = xid.New().String()
	db.db.Create(&choice)
}

func (db DB) FindChoice(choiceId string) models.Choice {
	var choice models.Choice
	db.db.Where(&models.Choice{Id: choiceId}).Find(&choice)
	return choice
}
