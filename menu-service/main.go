package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chandraawardhana01/digitalent-microservices/menu-service/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/add-menu", http.HandlerFunc(handler.AddMenu))

	fmt.Println("Menu service listen on port :8000")
	log.Panic(http.ListenAndServe(":8000", router))
}
