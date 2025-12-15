package health

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type ResponseHealth struct {
	*http.Response
}

func GetWebsiteHealth(url string) (*http.Response, time.Duration, error) {
	uri, err := CheckUrl(url)
	if err != nil {
		return nil, 0, fmt.Errorf("Error: Wrong URL. More details: %w", err)
	}

	start := time.Now()
	resp, err := http.Get(uri.String())
	if err != nil {
		return nil, 0, fmt.Errorf("Request - %w", err)
	}
	defer resp.Body.Close()
	timeTaken := time.Since(start)

	return resp, timeTaken, err

}

func CheckUrl(uri string) (*url.URL, error) {

	return url.ParseRequestURI(uri)
}
