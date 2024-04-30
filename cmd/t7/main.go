package main

import "fmt"

func main() {
	fmt.Println("Hello world! Pointers")
	// To create Pointer use use * syntax
	// var p *int32 // Will store pointer or memory address; default `nil`
	var pp *int32 = new(int32)
	*pp = 100 // Set the value at memory location pp is pointing to to 100
	// *-> Use to refrence the value of pointer
	var i int32 // Will store value; default `0`
	fmt.Printf("The value p points to %v", *pp)
	fmt.Printf("The value if i is %v", i)
	// & -> Get the memory address of the variable not the value
	pp = &i // pp refrences to the memory address of i so pp and i refrence to the same address
	pointersAndFunction()
}

func pointersAndFunction() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\n The memory location of the thing1 array is: %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\n The result is %v", result)
	fmt.Printf("\n The value of thing1 is %v", result)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\n The memory location of the thing2 array is: %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
