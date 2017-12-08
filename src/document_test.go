package main

import "testing"

func TestDocumentSetStatusAdd(t *testing.T) {
	var d document
	err := d.setStatus("add")
	if err != nil {
		t.Errorf("Expected setStatus to not return error")
	}

	if d.Blacklisted != true || d.UpdatedAt.IsZero() {
		t.Errorf("Expected to return status equals true")
	}
}

func TestDocumentSetStatusRemove(t *testing.T) {
	var d document
	err := d.setStatus("remove")
	if err != nil {
		t.Errorf("Expected setStatus to not return error")
	}

	if d.Blacklisted != false || d.UpdatedAt.IsZero() {
		t.Errorf("Expected to return status equals true")
	}
}

func TestDocumentSetStatusInvalid(t *testing.T) {
	var d document
	err := d.setStatus("xablau")
	if err == nil {
		t.Errorf("Expected setStatus to not return error")
	}
}
