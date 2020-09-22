package summpy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/chuross/taisho/internal/app/ext/ext_google"
	"golang.org/x/xerrors"
)

func Get(text string, sentLimit int) (*SummpyResult, error) {
	apiURL := os.Getenv("TAISHO_SUMMPY_URL")

	idToken, err := ext_google.GetIdToken(apiURL)
	if err != nil {
		return nil, xerrors.Errorf("id token get error: %w", err)
	}

	url, err := url.Parse(apiURL)
	if err != nil {
		return nil, xerrors.Errorf("summpy init parse error: %w", err)
	}
	url.Query().Set("text", text)
	url.Query().Set("sent_limit", string(sentLimit))

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, xerrors.Errorf("summpy request error: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", idToken))

	res, err := http.DefaultClient.Do(req)
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
