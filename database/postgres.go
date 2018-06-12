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
	db.AutoMigrate(&models.Person{}, &models.Group{})
	return DB{db}, nil;
}

func (db DB) AddPerson(person models.Person) {
	person.Id = xid.New().String()
	db.db.Create(&person)
}

func (db DB) AddPersonToGroup(memberId string, groupId string) {

	//update people
	user := models.Person{
		Id: memberId,
	}
	person := db.FindPersons(user)
	personGroups := append(person[0].Groups, groupId)
	db.db.Model(&user).Updates(models.Person{Groups: personGroups})

	//update group
	groupDetails := db.FindGroup(groupId)
	persons := append(groupDetails.Members, memberId)
	db.db.Model(&models.Group{Id: groupId,}).Updates(models.Group{Members: persons})
}

func (db DB) FindPersons(person models.Person) []models.Person {
	var results []models.Person
	db.db.Where(&models.Person{FirstName: person.FirstName, LastName: person.LastName}).Find(&results)
	return results
}

func (db DB) AddGroup(group models.Group) {
	group.Id = xid.New().String()
	db.db.Create(&group)
}

func (db DB) FindGroup(groupId string)models.Group {
	var group models.Group
	db.db.Where(&models.Group{Id: groupId}).Find(&group)
	return group
}

func (db DB) FindMembers(groupId string) []models.Person {
	var group models.Group
	var results []models.Person
	var personIds []string
	db.db.Where(models.Group{Id: groupId}).Find(&group)
	for _, id := range group.Members {
		personIds = append(personIds, id)
	}
	db.db.Where("id in (?)", personIds).Find(&results)
	return results
}
