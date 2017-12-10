package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// Constants to be used on validator
// CpfSize and CnpjSize are used on validator.
// Cpf and Cnpj strings are returned on validator context
const (
	Cpf     = "CPF"
	CpfSize = 11

	Cnpj     = "CNPJ"
	CnpjSize = 14
)

// validatorContext are returned after validate document.
// This contains info to be returned on validator output
type validatorContext struct {
	input   string
	variety string
	size    int
	mask    []int
	numbers []int
	digits  []int
}

// filter will return only valid numbers in string
func (v *validatorContext) filter() {
	reg, _ := regexp.Compile("[^0-9]+")
	v.input = reg.ReplaceAllString(v.input, "")
}

// setVarietySizeAndMask identify document by size and set variety, size and mask
func (v *validatorContext) setVarietySizeAndMask() error {
	switch size := len(v.input); size {
	case CpfSize:
		v.variety = Cpf
		v.size = CpfSize
		v.mask = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}

	case CnpjSize:
		v.variety = Cnpj
		v.size = CnpjSize
		v.mask = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	default:
		return errors.New("Invalid input size")
	}

	return nil
}

// calculateNumberAndDigits checks the number digits from document
func (v *validatorContext) calculateNumbersAndDigits() error {
	totalSum := 0
	baseIndex := v.size - 2

	// iterate over input and generate the numbers to sum
	for i, char := range v.input {
		strNum := fmt.Sprintf("%c", char)
		intNum, conversionErr := strconv.Atoi(strNum)
		if conversionErr != nil {
			return errors.New("Invalid conversion number")
		}

		// calculate input item * mask
		if i < baseIndex {
			v.numbers = append(v.numbers, intNum)
			totalSum += intNum * v.mask[i]
		}

		// calculate numbers x masks withous digits
		if i == baseIndex {
			module := totalSum % 11

			if module < 2 {
				v.digits = append(v.digits, 0)
			} else {
				v.digits = append(v.digits, (11 - module))
			}

			// validate calculated 1st digit with input
			if v.digits[0] != intNum {
				return errors.New("Invalid number")
			}
		}
		// add 1st digit to number list
		if i >= baseIndex {
			v.numbers = append(v.numbers, intNum)
		}
	}

	// add new item on mask to sum all numbers in list again
	index := v.mask[0] + 1
	v.mask = append([]int{index}, v.mask...)
	totalSum = 0
	baseIndex++

	// sum all numbers, including 1st digit
	for i, num := range v.numbers {
		if i < baseIndex {
			totalSum += num * v.mask[i]
		}
	}

	// calculate module for 2nd digit
	module := totalSum % 11
	if module < 2 {
		v.digits = append(v.digits, 0)
	} else {
		v.digits = append(v.digits, (11 - module))
	}

	// validate 2nd digit with input
	if v.digits[1] != v.numbers[baseIndex] {
		return errors.New("Invalid number")
	}

	// no errors
	return nil
}

// validate receives a string representation for document number
// Support validation for brazilian cpf or cnpj
func validate(input string) (validatorContext, error) {
	c := validatorContext{}
	c.input = input
	c.filter()

	sizeErr := c.setVarietySizeAndMask()
	if sizeErr != nil {
		return c, sizeErr
	}

	digitErr := c.calculateNumbersAndDigits()
	if digitErr != nil {
		return c, digitErr
	}

	return c, nil
}
