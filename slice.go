package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	numbersSlice := make([]int, 3, 3)
	var input string
	fmt.Println("Enter the numbers")
	index := 0
	for {

		fmt.Scan(&input)
		if input == "X" {
			break
		}
		integerValue, err := strconv.Atoi(input)
		if err != nil {
			return
		}
		if index < 3 {
			numbersSlice[0] = integerValue
			sort.Ints(numbersSlice)
		} else {
			numbersSlice = append(numbersSlice, integerValue)
			sort.Ints(numbersSlice)
		}

		for _, value := range numbersSlice {
			fmt.Print(value)
			fmt.Print(" ")
		}
		fmt.Println("")
		index++

	}
	sort.Ints(numbersSlice)
	for i := 0; i < len(numbersSlice); {
		fmt.Print(numbersSlice[i])
		fmt.Print(" ")
		i++
	}
	fmt.Println("")

}
