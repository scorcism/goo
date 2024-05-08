package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(("Hello world!"))
	// Function and control structure
	printMe()

	var num1 int = 11
	var num2 int = 2
	var res = intDivision(num1, num2)
	fmt.Println(res)
	var result, remainder, err = intDivisionRemainter(num1, num2)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the division is %v", result)
	} else {
		fmt.Printf("The result of the division is %v and remainder %v", result, remainder)
	}
}

func printMe() {
	fmt.Println("Hello world from printMe")
}

// - arg 1			arg 2			- function return type
func intDivision(numerator int, denominatior int) int {
	var result int = numerator / denominatior
	return result
}

// Returns multiple
func intDivisionRemainter(numerator int, denominator int) (int, int, error) {
	var err error // Error type; default value is `nil`
	if denominator == 0 {
		err = errors.New("Cannot divide by 0")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
