package main

import "fmt"

// Structs type typename struct
type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

type owner struct {
	name string
}

// Asigining fucntion to the struct
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type electricEngine struct {
	a uint8
	b uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.a * e.b
}

// Interface
type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("You cann't make it there!")
	}
}

func main() {

	fmt.Println("Hello world!")
	// Struct and Interfaces
	// var myEnginre gasEngine = gasEngine{mpg: 0, gallons: 0, owner{"ALex"}}
	var myEngine gasEngine = gasEngine{12, 23, owner{"Alex"}}
	fmt.Println(myEngine.gallons, myEngine.mpg, myEngine.ownerInfo.name)

	// Anonymouse struct
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 25}
	fmt.Println(myEngine2.gallons, myEngine2.mpg)

	// Struct with function
	fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())

	canMakeIt(myEngine, 50)

}
