package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world!")

	var strSlice = []string{"l", "a", "d", "d", "o"}
	// Strings
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var laddoStr = strBuilder.String()
	fmt.Printf("\n%v", laddoStr)
}
