package summpy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/xerrors"
)

func Get(text string, sentLimit int) (*SummpyResult, error) {
	url, err := url.Parse(os.Getenv("TAISHO_SUMMPY_URL"))
	if err != nil {
		return nil, xerrors.Errorf("summpy init parse error: %w", err)
	}
	url.Query().Set("text", text)
	url.Query().Set("sent_limit", string(sentLimit))

	res, err := http.Get(url.String())
	if err != nil {
		return nil, xerrors.Errorf("summpy request error: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, xerrors.Errorf("summpy request error: status=%d", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, xerrors.Errorf("summpy read body error: %w", err)
	}

	result := SummpyResult{}
	json.Unmarshal(b, &result)

	return &result, nil
}
