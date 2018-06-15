package database

import "github.com/suneetha1512/let-me-choose/models"

type Database interface {
	AddPerson(person models.Person)
	UpdatePersonChoice(memberId string, choiceId string)
	FindPersons(person models.Person) []models.Person
	FindPersonById(memberId string) models.Person
	AddChoice(choice models.Choice)
	FindChoice(choiceId string) models.Choice
} 
