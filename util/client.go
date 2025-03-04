package util

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/weather-app/monad"
)

// Pure
func BuildRequestURLWithAPIKey(baseURL string, params map[string]string, apiKey string) string {
	queryParams := url.Values{}
	for key, value := range params {
		queryParams.Add(key, value)
	}
	queryParams.Add("key", apiKey)
	return baseURL + "?" + queryParams.Encode()
}

// Impure
func FetchData[T any](url string) monad.IO[T] {
	return monad.IO[T]{Run: func() (T, error) {
		var target T
		resp, err := http.Get(url)
		if err != nil {
			return target, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return target, err
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return target, err
		}

		if err := json.Unmarshal(body, &target); err != nil {
			return target, err
		}

		return target, nil
	}}
}
