package main

import "testing"

// test validator
func TestValidatorNumberIsValid(t *testing.T) {
	cpf := "00549321063"

	isValid := validate(cpf)
	expected := true

	if isValid != expected {
		t.Errorf("Expected to return true for valid document number")
	}
}
