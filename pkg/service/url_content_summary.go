package service

import (
	"net/http"

	goose "github.com/advancedlogic/GoOse"
	"github.com/chuross/taisho/internal/app/ext/summpy.go"
	"github.com/chuross/taisho/pkg/model/url_content"
	"golang.org/x/xerrors"
)

func GetUrlContentSummary(url string) (*url_content.Summary, error) {
	ext := goose.New()
	res, err := http.Get(url)
	if err != nil {
		return nil, xerrors.Errorf("article fetch error: %w", err)
	}

	defer res.Body.Close()

	switch res.Header.Get("Content-Type") {
	case "text/plain":
		fallthrough
	case "text/html":
	default:
		return nil, nil
	}

	article, err := ext.ExtractFromURL(url)
	if err != nil {
		return nil, xerrors.Errorf("extract aricle error: %w", err)
	}

	_, err = summpy.Get(article.CleanedText, 3)
	if err != nil {
		return nil, xerrors.Errorf("summpy error: %w", err)
	}

	return &url_content.Summary{
		TopImageUrl: &article.TopImage,
		Summaries:   make([]string, 0),
	}, nil
}
