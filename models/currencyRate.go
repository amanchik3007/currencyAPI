package models

import (
	"testRates/repository"
	"time"
)

type Items struct {
	Id    int       `json:"id" db:"id"`
	Title string    `json:"title" db:"title"`
	Code  string    `json:"code" db:"code"`
	Value float32   `json:"value" db:"value"`
	ADate time.Time `json:"a_date" db:"a_date"`
}

func (item Items) Save() error {
	_, err := repository.DB.NamedExec(`INSERT INTO dbo.R_CURRENCY(title, code, value, a_date) VALUES (:title,:code,:value, :a_date)`,
		map[string]interface{}{
			"title":  item.Title,
			"code":   item.Code,
			"value":  item.Value,
			"a_date": item.ADate,
		})
	if err != nil {
		return err
	}
	return nil
}

func GetItemsByDate(date time.Time) []Items {
	items := make([]Items, 0)
	query := `SELECT * FROM dbo.R_CURRENCY WHERE a_date = $1`
	err := repository.DB.Select(&items, query, date)
	if err != nil {
		return []Items{}
	}
	return items
}

func GetItemsByDateAndCode(date time.Time, code string) []Items {
	items := make([]Items, 0)
	query := `SELECT * FROM dbo.R_CURRENCY WHERE a_date = $1 AND code = $2`
	err := repository.DB.Select(&items, query, date, code)
	if err != nil {
		return []Items{}
	}
	return items
}
