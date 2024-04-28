package main

import "fmt"

func main() {
	fmt.Println(("Hello world!"))

	// Array
	// var intArr [3]int32
	// var intArr2 [3]int32 = [3]int32{1, 2, 3}
	// intArr3 := [3]int32{1, 2, 3}
	// intArr4 := [...]int32{1, 2, 3}

	// Slice  loops

	// Map
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"A": 12, "1": 2}
	fmt.Println(myMap2["A"])
	// var age, ok = myMap2["Apple"] // ok will be false here means the key is not in mymap2
	// Dlete value from map
	// delete(myMap2, "A")

	// Loop - no order is preserverd
	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}

	// While in go; go ofically dont have while loop
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}
	for i:=0; i< 10; i++{
		fmt.Println(i)
	}
}
