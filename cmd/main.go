package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/suneetha1512/krick-ninjas/api"
)

func main() {
	r := httprouter.New()

	r.GET("/health", api.Health())

	fmt.Println("Starting service on 8432 port...")
	log.Fatal(http.ListenAndServe(":8432", r))
}
