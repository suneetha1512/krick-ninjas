package database

import "github.com/suneetha1512/let-me-choose/models"

type Database interface {
	AddPerson(person models.Person)
	FindPersons(person models.Person) []models.Person
	AddGroup(group models.Group)
	FindMembers(groupId string) []models.Person
	AddPersonToGroup(memberId string, groupId string)
} 
