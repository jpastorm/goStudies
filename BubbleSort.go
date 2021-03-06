package main

import (
	"fmt"
)

func bubbleSort(input [10]int) {
	// n is the number of items in our list
	n := 10
	// set swapped to true
	swapped := true
	// loop
	for swapped {
		// set swapped to false
		swapped = false
		// iterate through all of the elements in our list
		for i := 1; i < n; i++ {
			// if the current element is greater than the next
			// element, swap them
			if input[i-1] > input[i] {
				swap(&input[i], &input[i-1])
				// set swapped to true - this is important
				// if the loop ends and swapped is still equal
				// to false, our algorithm will assume the list is
				// fully sorted.
				swapped = true
			}
		}
	}
	// finally, print out the sorted list
	fmt.Println(input)
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	fmt.Println("Enter 10 integers")
	var vals [10]int
	i := 0
	for i < 10 {
		fmt.Scan(&vals[i])
		i++

	}

	bubbleSort(vals)
}
