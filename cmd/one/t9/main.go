package main

import (
	"fmt"
	
)

func main() {
	fmt.Println("Hello world! CHANNELS")

	// channels -> enables go routines to pass around information
	// Features of go rutines
	// Hold data
	// Thread safe
	// Listen for Data
	// Creating a channel
	// Using make function followd by chan key workd then the type of value the channel hold
	// var c = make(chan int)
	// // Add value 1 to channel
	// c <- 1
	// // retrieve value from the channle
	// var i = <- c

	// Un buffered channel
	var c = make(chan int)
	go processes(c)
	for i := range c {
		fmt.Println(i)
	}
	main2()

}

func processes(c chan int) {
	defer close(c) // Close the channel at end
	for i := 0; i < 5; i++ {
		c <- i
	}
}
