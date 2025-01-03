package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("--- BUBBLE SORT ---\n\n")

	unsortedList := []int{23, 4, 67, -8, 21}
	fmt.Printf("unsorted list: %v\n\n", unsortedList)

	startTime := time.Now()
	sortedList := bubbleSort(unsortedList)
	duration := time.Since(startTime)

	fmt.Printf("\nsorted list: %v\n", sortedList)
	fmt.Printf("execution time: %v\n", duration)
}

func bubbleSort(list []int) []int {
	iteration := 1 // aux variable for counting iterations

	for {
		// flag to check if slice is sorted
		isSorted := true

		for i := 0; i < len(list)-1; i++ {
			// swap position
			if list[i] > list[i+1] {
				list[i+1], list[i] = list[i], list[i+1]

				// array is not sorted
				isSorted = false
			}
		}

		// Print array after swapping
		fmt.Printf("%d iteration: %v\n", iteration, list)
		iteration++

		// break loop when array sorted
		if isSorted {
			break
		}
	}
	return list
}
