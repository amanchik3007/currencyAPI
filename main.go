package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"testRates/config"
	"testRates/repository"
	"testRates/routes"
)

func main() {
	fmt.Println("Starting the application...")
	configuration, err := config.LoadConfiguration("config.json")
	if err != nil{
		panic(err)
	}

	repository.OpenDB(configuration.Database.ConnectionString)
	defer repository.CloseDB()

	router := mux.NewRouter()

	routes.Handlers(router)
	fmt.Printf("Server running on port %s", configuration.Port)
	log.Fatal(http.ListenAndServe(":"+configuration.Port, router))
}