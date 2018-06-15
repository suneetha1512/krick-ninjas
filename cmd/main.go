package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/suneetha1512/let-me-choose/api"
	"github.com/suneetha1512/let-me-choose/database"
)

func main() {

	// initialize DB
	fmt.Println("Setting up database..")
	db, err := database.Setup("user=user password=pass host=172.16.200.32 dbname=ninjas sslmode=disable")
	if err != nil {
		log.Fatal("failed to create DB conn", err)
	}

	r := httprouter.New()

	r.GET("/health", api.Health())

	r.POST("/person", api.AddPerson(db))
	r.GET("/person", api.GetPerson(db))
	r.POST("/choice", api.AddChoice(db))
	r.POST("/choice/:choiceId/member/:memberId", api.AddMemberToChoice(db))

	fmt.Println("Starting service on 8432 port...")
	log.Fatal(http.ListenAndServe(":8432", r))
}
