package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/yukihir0/gec"

	"github.com/chuross/taisho/internal/ext/summpy.go"
	"github.com/chuross/taisho/pkg/model/url_content"
	"github.com/chuross/taisho/pkg/util"
	"golang.org/x/xerrors"
)

var (
	ignoreHosts = []string{
		"twitter.com",
		"amazon.co.jp",
	}
)

func GetURLContentSummary(url string) (*url_content.Summary, error) {
	for _, host := range ignoreHosts {
		if strings.Contains(url, host) {
			return &url_content.Summary{
				Summaries: make([]string, 0),
			}, nil
		}
	}

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

	bodyStr, err := util.ToUTF8(body)
	if err != nil {
		return nil, xerrors.Errorf("content to utf8 error: %w", err)
	}

	content, _ := gec.Analyse(bodyStr, nil)
	if err != nil {
		return nil, xerrors.Errorf("content extract error: %w", err)
	}

	if len(content) == 0 {
		log.Printf("content is empty")
		return nil, nil
	}

	result, err := summpy.Get(content, 4)
	if err != nil {
		return nil, xerrors.Errorf("summpy error: %w", err)
	}

	return &url_content.Summary{
		Summaries: result.Summaries,
	}, nil
}
