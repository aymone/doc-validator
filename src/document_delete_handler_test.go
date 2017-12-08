// +build acceptance

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestDocumentDeleteNotFound(t *testing.T) {
	httpClient := &http.Client{
		Timeout: 1 * time.Second,
	}

	id := "18243547380"
	url, err := url.Parse(fmt.Sprintf("%v/documents/%s", host, id))
	if err != nil {
		log.Fatal(err)
	}

	req := &http.Request{
		Method: "DELETE",
		URL:    url,
	}

	res, err := httpClient.Do(req)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 404 {
		t.Errorf("Expected 404 NOT FOUND, got %v", res.StatusCode)
	}
}

func TestDocumentCreateForDelete(t *testing.T) {
	httpClient := &http.Client{
		Timeout: 1 * time.Second,
	}

	url := fmt.Sprintf("%v/documents", host)
	doc := document{ID: "18243547380"}

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

func TestDocumentDelete(t *testing.T) {
	httpClient := &http.Client{
		Timeout: 1 * time.Second,
	}

	id := "18243547380"
	url, err := url.Parse(fmt.Sprintf("%v/documents/%s", host, id))
	if err != nil {
		log.Fatal(err)
	}

	req := &http.Request{
		Method: "DELETE",
		URL:    url,
	}

	res, err := httpClient.Do(req)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 204 {
		t.Errorf("Expected 204 NO CONTENT, got %v", res.StatusCode)
	}
}
