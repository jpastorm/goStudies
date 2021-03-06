package main

import (
	"fmt"
	"math"
	"sort"
)

func Sort(slice []int, channels chan []int) {
	fmt.Printf("slice: %v\n", slice)
	sort.Ints(slice)
	channels <- slice
}

func main() {
	inputs := make([]int, 0)
	input := 0
	fmt.Printf("enter digits, or any letter to exit.\n> ")
	n, e := fmt.Scan(&input)
	for n > 0 || e == nil {
		inputs = append(inputs, input)
		fmt.Println(inputs)
		fmt.Printf("> ")
		n, e = fmt.Scan(&input)
	}
	length := len(inputs)
	if length == 0 {
		fmt.Printf("bye.\n\n")
		return
	}
	fmt.Println()

	splitSize := 4
	channels := make([]chan []int, 0)
	channelsIndex := 0
	for i, c, l, slice := splitSize, 0, length, inputs[:]; l > 0; i-- {
		c = int(math.Ceil(float64(l) / float64(i)))
		channels = append(channels, make(chan []int))
		channelsIndex++
		slice = slice[:c]
		go Sort(slice, channels[channelsIndex-1])
		slice = slice[c:c]
		l = l - c
	}
	sortedSlices := make([][]int, splitSize)
	for i := 0; i < channelsIndex; i++ {
		sortedSlices[i] = <-channels[i]
	}

	result := make([]int, 0)
	sliceIndex := make([]int, splitSize)
	tmp := 0
	for j := 0; j < length; j++ {
		x := 0
		for i := 0; i < splitSize; i++ {
			if len(sortedSlices[i]) > sliceIndex[i] {
				tmp = sortedSlices[i][sliceIndex[i]]
				x = i
				break
			}
		}
		for i := 0; i < splitSize; i++ {
			if len(sortedSlices[i]) > sliceIndex[i] {
				if tmp > sortedSlices[i][sliceIndex[i]] {
					tmp = sortedSlices[i][sliceIndex[i]]
					x = i
				}
			}
		}
		sliceIndex[x]++
		result = append(result, tmp)
	}
	fmt.Printf("\nsorted array: %v\n\n", result)
}
