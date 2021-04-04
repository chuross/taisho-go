package summpy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	google_ext "github.com/chuross/taisho/internal/ext/google"
	"golang.org/x/xerrors"
)

func Get(text string, sentLimit int) (*SummpyResult, error) {
	apiURL := os.Getenv("TAISHO_SUMMPY_URL")

	idToken, err := google_ext.GetIdToken(apiURL)
	if err != nil {
		return nil, xerrors.Errorf("id token get error: %w", err)
	}

	values := url.Values{}
	values.Add("text", text)
	values.Add("sent_limit", strconv.Itoa(sentLimit))

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, xerrors.Errorf("summpy request error: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", idToken))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
