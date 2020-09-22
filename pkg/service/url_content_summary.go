package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/yukihir0/gec"

	"github.com/chuross/taisho/internal/app/ext/summpy.go"
	"github.com/chuross/taisho/pkg/model/url_content"
	"golang.org/x/xerrors"
)

func GetUrlContentSummary(url string) (*url_content.Summary, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, xerrors.Errorf("content fetch error: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, xerrors.Errorf("content fetch error: status=%d", res.StatusCode)
	}

	switch {
	case strings.Contains(res.Header.Get("Content-Type"), "text/plain"):
		fallthrough
	case strings.Contains(res.Header.Get("Content-Type"), "text/html"):
	default:
		return &url_content.Summary{
			Summaries: make([]string, 0),
		}, nil
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, xerrors.Errorf("content read error: %w", err)
	}

	content, _ := gec.Analyse(string(body), nil)
	if err != nil {
		return nil, xerrors.Errorf("content extract error: %w", err)
	}

	if len(content) == 0 {
		log.Printf("content is empty")
		return nil, nil
	}

	result, err := summpy.Get(content, 3)
	if err != nil {
		return nil, xerrors.Errorf("summpy error: %w", err)
	}

	return &url_content.Summary{
		Summaries: result.Summaries,
	}, nil
}
