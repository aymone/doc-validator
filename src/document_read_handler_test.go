// +build acceptance

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestDocumentReadNotFound(t *testing.T) {
	httpClient := &http.Client{
		Timeout: 1 * time.Second,
	}

	id := "46616939731"
	url := fmt.Sprintf("%v/documents/%s", host, id)
	res, err := httpClient.Get(url)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 404 {
		t.Errorf("Expected 404 NOT FOUND, got %v", res.StatusCode)
	}
}

func TestDocumentCreateForRead(t *testing.T) {
	httpClient := &http.Client{
		Timeout: 1 * time.Second,
	}

	url := fmt.Sprintf("%v/documents", host)
	doc := document{ID: "46616939731"}

	body, _ := json.Marshal(doc)

	res, err := httpClient.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("Expected 200 OK, got %v", res.StatusCode)
	}
}

func TestDocumentRead(t *testing.T) {
	httpClient := &http.Client{
		Timeout: 1 * time.Second,
	}

	id := "46616939731"
	url := fmt.Sprintf("%v/documents/%s", host, id)
	res, err := httpClient.Get(url)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("Expected 200 OK, got %v", res.StatusCode)
	}
}
