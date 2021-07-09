package models

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type Rates struct {
	Generator   string     `xml:"generator"`
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	Copyright   string     `xml:"copyright"`
	Date        string     `xml:"date"`
	Items       []ItemsXml `xml:"item"`
}

type ItemsXml struct {
	FullName    string  `xml:"fullname"`
	Title       string  `xml:"title"`
	Description float32 `xml:"description"`
	Quant       int     `xml:"quant"`
	Index       string  `xml:"index"`
	Change      string  `xml:"change"`
}

func GetCurrency(date string) (Rates, error) {
	var rates Rates
	resp, err := http.Get("https://nationalbank.kz/rss/get_rates.cfm?fdate=" + date)
	if err != nil {
		return Rates{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Rates{}, err
	}

	err = xml.Unmarshal(body, &rates)
	if err != nil {
		return Rates{}, err
	}
	return rates, nil
}
