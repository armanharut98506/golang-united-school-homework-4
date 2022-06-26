package main

import (
	"fmt"
	"strconv"
	"regexp"
	"errors"
) 

var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	errorNotValidInput = errors.New("error is not valid")
)

func StringSum(input string) (output string, err error) {
	if input == "" {
		return "", errorEmptyInput
	}

	reNums := regexp.MustCompile("(\\+|-)?\\d+")	
	nums := reNums.FindAllString(input, -1)

	reNotValid := regexp.MustCompile("[^(\\d|\\+|\\-|\\W)]")
	notValid := reNotValid.FindAllString(input, -1)

	if len(nums) > 2 {
		return "", errorNotTwoOperands
	} else if len(notValid) > 0 {
		return "", errorNotValidInput
	}

	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])

	return strconv.Itoa(num1 + num2), nil
}

func main() {
	input := " -8  -5 "
	fmt.Println(StringSum(input))
}