package main

import (
	"fmt"
	"strings"
)

func main() {
	var stringValue string
	fmt.Println("Enter a string")
	fmt.Scan(&stringValue)
	stringValue = strings.ToLower(stringValue)
	if strings.HasPrefix(stringValue, "i") && strings.Contains(stringValue, "a") && strings.HasSuffix(stringValue, "n") {
		fmt.Println("Found!")
	} else {

		fmt.Println("Not Found!")
	}

}
