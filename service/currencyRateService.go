package service

import (
	"log"
	"testRates/models"
	"time"
)

func SaveCurrencyRateToDatabase(date string) {
	rates, err := models.GetCurrency(date)
	if err != nil {
		log.Print(err)
		return
	}

	aDate, err := time.Parse("02.01.2006", rates.Date)
	if err != nil {
		log.Print(err)
		return
	}

	for _, v := range rates.Items {
		item := models.Items{
			Title: v.FullName,
			Code:  v.Title,
			Value: v.Description,
			ADate: aDate,
		}
		err := item.Save()
		if err != nil {
			log.Print(err)
		}
	}
}

func GetCurrency(date, code string) []models.Items {
	aDate, err := time.Parse("02.01.2006", date)
	if err != nil {
		log.Print(err)
		return []models.Items{}
	}
	if date != "" && code != "" {
		items := models.GetItemsByDateAndCode(aDate, code)
		return items
	} else {
		items := models.GetItemsByDate(aDate)
		return items
	}
}
