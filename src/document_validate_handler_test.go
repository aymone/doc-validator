// +build acceptance

package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestDocumentValidateCpf(t *testing.T) {
	httpClient := &http.Client{
		Timeout: 1 * time.Second,
	}

	docID := "39390332516"
	url := fmt.Sprintf("%v/documents/%s/validate", host, docID)

	res, err := httpClient.Get(url)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected 200 OK, got %v", res.StatusCode)
	}
}

func TestDocumentValidateCnpj(t *testing.T) {
	httpClient := &http.Client{
		Timeout: 1 * time.Second,
	}

	docID := "10329388000122"
	url := fmt.Sprintf("%v/documents/%s/validate", host, docID)

	res, err := httpClient.Get(url)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected 200 OK, got %v", res.StatusCode)
	}
}
