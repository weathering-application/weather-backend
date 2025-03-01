package utils

import "net/url"

// Pure function
func BuildRequestURL(baseURL string, params map[string]string, apiKey string) string {
	queryParams := url.Values{}
	for key, value := range params {
		queryParams.Add(key, value)
	}
	queryParams.Add("key", apiKey)
	return baseURL + "?" + queryParams.Encode()
}
