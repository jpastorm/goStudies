package main

import "fmt"

func main() {
	var floatValue float32
	fmt.Println("Enter a floating point value")
	fmt.Scan(&floatValue)
	var intValue int32
	intValue = int32(floatValue)
	fmt.Println("The integer value is", intValue)
}

//
