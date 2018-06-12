package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/suneetha1512/let-me-choose/database"
	"github.com/suneetha1512/let-me-choose/models"
)

func AddPerson(db database.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		data := models.Person{}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			fmt.Printf("failed to decode person body %s", err)
			invalidRequest(w)
			return
		}
		db.AddPerson(data)
	}
}

func GetPerson(db database.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		firstName := r.URL.Query().Get("first")
		lastName := r.URL.Query().Get("last")
		person := models.Person{FirstName: firstName, LastName: lastName}
		results := db.FindPersons(person)
		successResponse(w, results)
	}
}

func AddGroup(db database.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		data := models.Group{}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			fmt.Printf("failed to decode group body %s", err)
			invalidRequest(w)
			return
		}
		db.AddGroup(data)
	}
}

func AddMemberToGroup(db database.DB) httprouter.Handle  {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		groupId := p.ByName("groupId")
		memberId := p.ByName("memberId")
		db.AddPersonToGroup(memberId, groupId)
	}
}

func FindMembers(db database.DB) httprouter.Handle  {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		groupId := p.ByName("groupId")
		results := db.FindMembers(groupId)
		successResponse(w, results)
	}
}
