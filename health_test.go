package main

import (
	"testing"
)

func TestUrlPerfect(t *testing.T) {

	url := "https://www.google.com"

	want := &Response{
		Code:    200,
		Message: url + " is healthy!",
	}
	resp := CheckHealth(url)
	if want.Message != resp.Message ||
		want.Code != resp.Code {
		t.Errorf("Healthy URL Test failed\nCode: %d,\nMessage: %q", resp.Code, resp.Message)
	}
}

func TestUrlBad(t *testing.T) {
	url := "https://www.gdssadoogle.com"

	resp := CheckHealth(url)
	if resp.Code != 0 {
		t.Errorf("Wrong URL Test failed\nCode: %d,\nMessage: %q", resp.Code, resp.Message)
	}
}
