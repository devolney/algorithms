package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("--- INSERTION SORT ---\n\n")

	unsortedList := []int{23, 4, 67, -8, 21}
	fmt.Printf("unsorted list: %v\n\n", unsortedList)

	startTime := time.Now()
	sortedList := insertionSort(unsortedList)
	duration := time.Since(startTime)

	fmt.Printf("\nsorted list: %v\n", sortedList)
	fmt.Printf("execution time: %v\n", duration)
}

func insertionSort(list []int) []int {
	// begin in index 1
	for i := 1; i < len(list); i++ {
		key := list[i]
		j := i - 1

		// Move elements bigger than 'key' one position forward
		for j >= 0 && list[j] > key {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = key // Insert 'key' in correct position

		fmt.Printf("%dª iteração: %v\n", i+1, list)
	}

	return list
}
