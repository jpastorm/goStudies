package main

import (
	"encoding/json"
	"fmt"
)

type testType struct {
	newValue interface{}
}

func main() {
	var a testType
	var mongoChangeEvent testType
	a.newValue = 318800000000
	err := json.Unmarshal([]byte(msg.Value), &mongoChangeEvent)
	str := fmt.Sprintf("%v", a)
	fmt.Printf(str)
}
