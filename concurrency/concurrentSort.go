package main

import (
	"fmt"
)

func bubbleSort(input []int, channel chan []int) {
	n := len(input)
	swapped := true

	for swapped {

		swapped = false

		for i := 1; i < n; i++ {

			if input[i-1] > input[i] {
				swap(&input[i], &input[i-1])
				swapped = true
			}
		}
	}
	// finally, print out the sorted list
	fmt.Println(input)
	channel <- input
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	fmt.Println("Enter 12 integers")
	var vals [12]int
	i := 0
	for i < 12 {
		fmt.Scan(&vals[i])
		i++

	}
	aChan := make(chan []int)
	bChan := make(chan []int)
	cChan := make(chan []int)
	dChan := make(chan []int)
	go bubbleSort(vals[:3], aChan)
	go bubbleSort(vals[3:6], bChan)
	go bubbleSort(vals[6:9], cChan)
	go bubbleSort(vals[9:], dChan)
	a := <-aChan
	b := <-bChan
	c := <-cChan
	d := <-dChan
	intermediateArray1 := merge(a, b)
	intermediateArray2 := merge(c, d)
	finalArray := merge(intermediateArray1, intermediateArray2)
	fmt.Println(finalArray)

}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}
