package routes

import (
	"github.com/gorilla/mux"
	"testRates/controllers"
)

func Handlers(router *mux.Router) {
	router.HandleFunc("/currency/save/{date}", controllers.SaveRates).Methods("GET")
	router.HandleFunc("/currency/{date}", controllers.GetRatesById).Methods("GET")
	router.HandleFunc("/currency/{date}/{code}", controllers.GetRatesById).Methods("GET")
}