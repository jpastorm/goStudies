package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string
	fmt.Println("Enter name")
	fmt.Scan(&name)
	fmt.Println("Enter an address")
	fmt.Scanln(&address)
	m := make(map[string]string)
	m["name"] = name
	m["address"] = address
	jsonMap, err := json.Marshal(m)
	if err != nil {
		return
	}
	fmt.Println(string(jsonMap))

}
