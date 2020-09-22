package summpy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	google_ext "github.com/chuross/taisho/internal/app/ext/google"
	"golang.org/x/xerrors"
)

func Get(text string, sentLimit int) (*SummpyResult, error) {
	apiURL := os.Getenv("TAISHO_SUMMPY_URL")

	idToken, err := google_ext.GetIdToken(apiURL)
	if err != nil {
		return nil, xerrors.Errorf("id token get error: %w", err)
	}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, xerrors.Errorf("summpy request error: %w", err)
	}

	q := req.URL.Query()
	q.Add("text", text)
	q.Add("sent_limit", strconv.Itoa(sentLimit))
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", idToken))

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
