package database

import "github.com/suneetha1512/krick-ninjas/models"

type Database interface {
	AddPerson(person models.Person)
	FindPersons(person models.Person) []models.Person
} 
