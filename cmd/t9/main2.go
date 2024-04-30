package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_PRICE float32 = 5

func main2() {
	fmt.Println("Main2 channels")
	var channel = make(chan string)
	var websites = []string{"A.com", "B.com", "C.com"}
	for i := range websites {
		go checkPrice(websites[i], channel)
	}
	sendMessage(channel)
}

func checkPrice(website string, channel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var checkPrice = rand.Float32() * 20
		if checkPrice <= MAX_PRICE {
			channel <- website
			break
		}
	}
}

func sendMessage(channel chan string) {
	fmt.Printf("\nFound a deal on at %s", <-channel)
}
