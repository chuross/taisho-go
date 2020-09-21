package service

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/xerrors"
)

const weatherUrl = "https://weather.yahoo.co.jp/weather/jp/raincloud/3.html"

func GetWeatherImage() (*string, error) {
	doc, err := goquery.NewDocument(weatherUrl)
	if err != nil {
		return nil, xerrors.Errorf("get weather node error: %w", err)
	}

	elm := doc.Find("#imgDatCh .mainImg img")
	if elm == nil {
		return nil, xerrors.New("element not found")
	}

	url, _ := elm.First().Attr("src")
	return &url, nil
}
