package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("--- MERGE SORT ---\n\n")

	unsortedList := []int{23, 4, 67, -8, 21}
	fmt.Printf("unsorted list: %v\n\n", unsortedList)

	startTime := time.Now()
	sortedList := mergeSort(unsortedList)
	duration := time.Since(startTime)

	fmt.Printf("\nsorted list: %v\n", sortedList)
	fmt.Printf("execution time: %v\n", duration)
}

func mergeSort(list []int) []int {
	// base case
	if len(list) <= 1 {
		return list
	}

	// recursive case
	middle := len(list) / 2
	leftSlice := mergeSort(list[:middle])
	rightSlice := mergeSort(list[middle:])
	return merge(leftSlice, rightSlice)
}

func merge(left []int, right []int) []int {
	// Create a slice to store the merged result.
	result := make([]int, len(left)+len(right))

	/*
		idx: index variable for the result slice.
		i: index variable for the left slice.
		j: index variable for the right slice.
	*/
	var idx, i, j int

	// Compare elements from both slices and merge them in sorted order
	// until one of the slices is fully traversed.
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[idx] = left[i]
			i++
		} else {
			result[idx] = right[j]
			j++
		}
		idx++
	}

	//---
	// At this point, only one of the slices has remaining elements.

	// Add any remaining elements from the left slice to the result slice.
	for i < len(left) {
		result[idx] = left[i]
		idx++
		i++
	}

	// Add any remaining elements from the right slice to the result slice.
	for j < len(right) {
		result[idx] = right[j]
		idx++
		j++
	}

	return result
}
