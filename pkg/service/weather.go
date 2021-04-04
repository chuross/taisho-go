package service

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/xerrors"
)

const weatherURL = "https://tenki.jp/radar/"

func GetWeatherImage() (*string, error) {
	doc, err := goquery.NewDocument(weatherURL)
	if err != nil {
		return nil, xerrors.Errorf("get weather node error: %w", err)
	}

	elm := doc.Find("#radar-image")
	if elm == nil {
		return nil, xerrors.New("element not found")
	}

	url, _ := elm.First().Attr("src")
	return &url, nil
}
