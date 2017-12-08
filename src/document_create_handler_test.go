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

func TestDocumentCreateCpf(t *testing.T) {
	httpClient := &http.Client{
		Timeout: 1 * time.Second,
	}

	url := fmt.Sprintf("%v/documents", host)
	doc := document{ID: "005.493.210-63"}

	body, _ := json.Marshal(doc)

	res, err := httpClient.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("Expected 200, got %v", res.StatusCode)
	}
}

func TestDocumentCreateCnpj(t *testing.T) {
	httpClient := &http.Client{
		Timeout: 1 * time.Second,
	}

	url := fmt.Sprintf("%v/documents", host)
	doc := document{ID: "05.778.335/0001-59"}

	body, _ := json.Marshal(doc)

	res, err := httpClient.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("Expected 200, got %v", res.StatusCode)
	}
}

func TestDocumentDuplicatedError(t *testing.T) {
	httpClient := &http.Client{
		Timeout: 1 * time.Second,
	}

	url := fmt.Sprintf("%v/documents", host)
	doc := document{ID: "00549321063"}

	body, _ := json.Marshal(doc)

	res, err := httpClient.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 409 {
		t.Errorf("Expected 409, CONFLICT, got %v", res.StatusCode)
	}
}

func TestDocumentInvalidError(t *testing.T) {
	httpClient := &http.Client{
		Timeout: 1 * time.Second,
	}

	url := fmt.Sprintf("%v/documents", host)
	doc := document{ID: "00549321064"}

	body, _ := json.Marshal(doc)

	res, err := httpClient.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 400 {
		t.Errorf("Expected 400, BAD REQUEST, got %v", res.StatusCode)
	}
}
