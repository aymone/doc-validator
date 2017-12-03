package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const (
	Cpf     = "cpf"
	CpfSize = 11

	Cnpj     = "cnpj"
	CnpjSize = 14
)

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

func (v *validatorContext) calculateNumbersAndDigits() error {
	totalSum := 0
	baseIndex := v.size - 2

	for i, char := range v.input {
		strNum := fmt.Sprintf("%c", char)
		intNum, conversionErr := strconv.Atoi(strNum)
		if conversionErr != nil {
			return errors.New("Invalid conversion number")
		}

		if i < baseIndex {
			v.numbers = append(v.numbers, intNum)
			totalSum += intNum * v.mask[i]
		}

		if i == baseIndex {
			module := totalSum % 11

			if module < 2 {
				v.digits = append(v.digits, 0)
			} else {
				v.digits = append(v.digits, (11 - module))
			}

			if v.digits[0] != intNum {
				return errors.New("Invalid number")
			}
		}

		if i >= baseIndex {
			v.numbers = append(v.numbers, intNum)
		}
	}

	index := v.mask[0] + 1
	v.mask = append([]int{index}, v.mask...)
	totalSum = 0
	baseIndex++

	for i, num := range v.numbers {
		if i < baseIndex {
			totalSum += num * v.mask[i]
		}
	}

	module := totalSum % 11
	if module < 2 {
		v.digits = append(v.digits, 0)
	} else {
		v.digits = append(v.digits, (11 - module))
	}

	if v.digits[1] != v.numbers[baseIndex] {
		return errors.New("Invalid number")
	}

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
