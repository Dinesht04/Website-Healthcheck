package main

import (
	"net/http"
	"net/url"
)

type ResponseHealth struct {
	*http.Response
}

func GetWebsiteHealth(url string) (*http.Response, error) {
	uri, err := CheckUrl(url)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(uri.String())

	return resp, err

}

func CheckUrl(uri string) (*url.URL, error) {

	return url.ParseRequestURI(uri)
}
