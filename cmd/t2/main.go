package main

import "fmt"

func main() {
	var intNum int // To declare variable use var keyword
	fmt.Println(intNum)

	var stringA string = `Apple 
	in 
	next line
	`

	var stringB string = "In one line"

	fmt.Println(stringA)
	fmt.Println(stringB)

	// variable type infered
	var apple = "Apple" // Type infered as string

	myVar := "text" // shorthand
	fmt.Println(apple, myVar)

	const pi float32 = 3.1415
	fmt.Println(pi)
}
