/*
Package main implements a command-line interface for the Binary Search algorithm

It accepts a list of sorted integers followed by a target key as command-line arguments.
It then prints the index of the target key if found, or a message indicating it is missing.

Usage:

	go run binary_search_2_args.go [sorted numbers] [target key]

Example:

	go run binary_search_2_args.go 1 2 3 4 5 6 7 8 9 10 6
	Output: Target value is at position: 5

Note:

	The input list of numbers MUST be sorted for binary search to work correctly.
	If the list is not sorted, the results will be unpredictable.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run file name [sorted numbers] [key]")
	}

	fmt.Println("\t-----\tBINARY SEARCH\t-----")

	// Slice the args to separate the list form the key
	// We use len(os.Args)-1 to exclude the last element (the key) from the list
	fmt.Printf("The slice will be s = %v\n", os.Args[1:len(os.Args)-1])
	fmt.Printf("The slice type is: %T\n", os.Args[1:])
	fmt.Printf("The key is: %v\n", os.Args[len(os.Args)-1])
	fmt.Printf("Because s is a slice of strings we need to convert it to a slice of int\n")

	// Pre-allocate memory for the integer slice
	s := make([]int, 0, len(os.Args)-2)

	// Convert the last argument to the search key
	key, err := strconv.Atoi(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal("Failed to convert string to integer ", err)
	}

	fmt.Printf("key = %d\n", key)

	// Convert the rest of the arguments to the integer slice
	for _, arg := range os.Args[1 : len(os.Args)-1] {
		intArg, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatal("Failed to convert string to integer ", err)
		}

		s = append(s, intArg)
	}

	fmt.Printf("s = %v\n", s)

	// Perform the search
	targetValue := binarySearch(s, key)

	if targetValue == -1 {
		fmt.Println("The target value is not in the array")
	} else {
		fmt.Printf("The target value is at position: %d\n", targetValue)
	}
}

// binarySearch performs a logarithmic time O(log n) search on a sorted slice.
// It returns the index of the key if found, or -1 if not found
func binarySearch(s []int, key int) int {
	low := 0
	high := len(s) - 1

	for low <= high {
		// Calculate mid to avoid potential integer overflow
		med := low + (high-low)/2

		switch {
		case s[med] == key:
			return med
		case s[med] < key:
			low = med + 1
		default: // s[med] > key
			high = med - 1
		}
	}

	return -1
}
