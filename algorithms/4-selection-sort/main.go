package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("--- SELECTION SORT ---\n\n")

	unsortedList := []int{23, 4, 67, -8, 21}
	fmt.Printf("unsorted list: %v\n\n", unsortedList)

	startTime := time.Now()
	sortedList := selectionSort(unsortedList)
	duration := time.Since(startTime)

	fmt.Printf("\nsorted list: %v\n", sortedList)
	fmt.Printf("execution time: %v\n", duration)
}

func selectionSort(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		smallest := list[i] // aux varible to store smallest value
		var index int       // aux varible to store the smallest value's index

		// find smallest element
		for j := i + 1; j < len(list); j++ {
			if smallest > list[j] {
				smallest = list[j]
				index = j
			}
		}

		if smallest < list[i] {
			list[i], list[index] = list[index], list[i] // swap values

		}
		fmt.Printf("%d° iteration: %v\n", i+1, list)
	}

	return list
}
