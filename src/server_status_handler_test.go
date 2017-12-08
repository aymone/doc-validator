// +build acceptance

package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestServerStatus(t *testing.T) {
	httpClient := &http.Client{
		Timeout: 1 * time.Second,
	}

	url := fmt.Sprintf("%v/status", host)
	res, err := httpClient.Get(url)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("Expected 200 OK, got %v", res.StatusCode)
	}
}
