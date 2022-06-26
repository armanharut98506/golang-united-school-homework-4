package string_sum

import (
	"fmt"
	"errors"
	"regexp"
	"strconv"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	errorNotValidInput = errors.New("input not valid")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	if input == "" {
		return "", fmt.Errorf(errorEmptyInput.Error())
	}

	reNums := regexp.MustCompile("(\\+|-)?\\d+")	
	nums := reNums.FindAllString(input, -1)

	reNotValid := regexp.MustCompile("[a-zA-Z]")
	notValid := reNotValid.FindAllString(input, -1)

	if len(nums) != 2 {
		return "", fmt.Errorf(errorNotTwoOperands.Error())
	} else if len(notValid) > 0 {
		_, err := strconv.Atoi(notValid[0])
		err = fmt.Errorf(err.Error())
		return "", err
	}

	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])

	return strconv.Itoa(num1 + num2), nil
}
