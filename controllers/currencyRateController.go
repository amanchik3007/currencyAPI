package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"testRates/service"
	"testRates/utils"
)

func SaveRates(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]

	go service.SaveCurrencyRateToDatabase(date)

	utils.SuccessResponse(w, nil)
}

func GetRatesById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]
	code := params["code"]

	items := service.GetCurrency(date, code)

	utils.SuccessResponse(w, items)
}