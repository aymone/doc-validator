package main

import (
	"fmt"
	"testing"
)

func TestValidatorFilter(t *testing.T) {
	c := validatorContext{input: "005.493.210-63"}
	c.filter()

	if c.input != "00549321063" {
		t.Errorf("Expected to return cpf without non numeric chars")
	}
}

func TestValidatorSetVarietySizeForCpf(t *testing.T) {
	c := validatorContext{input: "005.493.210-63"}
	c.filter()
	c.setVarietySizeAndMask()

	if c.variety != Cpf {
		t.Errorf("Expected to return cpf in variety")
	}

	if c.size != CpfSize {
		t.Errorf("Expected to return cpf in variety")
	}
}

func TestValidatorSetVarietySizeForCnpj(t *testing.T) {
	c := validatorContext{input: "13.347.016/0001-17"}
	c.filter()
	c.setVarietySizeAndMask()

	if c.size != CnpjSize {
		t.Errorf("Expected to return cpf in variety")
	}
}

func TestValidatorWithInValidSize(t *testing.T) {
	invalidSize := "133470160"
	_, err := validate(invalidSize)
	if err == nil {
		t.Errorf("Expected to return errors")
	}
}

// test validator
func TestValidatorCpfIsValid(t *testing.T) {
	cpf := "00549321063"
	_, err := validate(cpf)
	if err != nil {
		fmt.Println(err)
		t.Errorf("Expected to not return errors")
	}
}

func TestValidatorCpfIsInValid(t *testing.T) {
	cpf := "00549321064"
	_, err := validate(cpf)
	if err == nil {
		t.Errorf("Expected to return errors")
	}
}

// test validator
func TestValidatorCnpjIsValid(t *testing.T) {
	cnpj := "13347016000117"
	_, err := validate(cnpj)
	if err != nil {
		fmt.Println(err)
		t.Errorf("Expected to not return errors")
	}
}

func TestValidatorCnpjIsInValid(t *testing.T) {
	cnpj := "13347016000118"
	_, err := validate(cnpj)
	if err == nil {
		t.Errorf("Expected to return errors")
	}
}
