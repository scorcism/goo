package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}
var wg = sync.WaitGroup{}
var mutex = sync.Mutex{}
var rwMutex = sync.RWMutex{}

func main() {
	fmt.Println(("Hello world! Goroutines"))

	t0 := time.Now()
	// for i := 0; i < len(dbData); i++ {
	// 	dbCall(i)
	// }
	for i := 0; i < len(dbData); i++ {
		// adding goroutines here by adding key `go`; But the program wont print anything has this will be an backend task
		// go dbCall(i)
		// To overcome the abouve one we can use wait group; Its like a counter
		wg.Add(1) // Adding one to the counter
		go dbCall(i)

	}
	wg.Wait() // Wait for the wg to go back to 0
	fmt.Printf("\nTotal execuation time: %v", time.Since(t0))
	fmt.Printf("\nVlaue of results: %v", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Microsecond)
	fmt.Println("The result from the database is: ", dbData[i])
	mutex.Lock()
	results = append(results, dbData[i]) // Lock the access as this will be an shared resourse, critical section
	mutex.Unlock()
	wg.Done() // Done to decrement the incremented counter
}

func save(result string){
	
}

