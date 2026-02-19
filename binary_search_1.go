package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	binarySearchInt(s, 6)
}

func binarySearchInt(s []int, key int) int {
	low := 0
	high := len(s) - 1

	for low <= high {
		mid := (high + low) / 2
		fmt.Printf("--- Mid is %d ---\n", mid)

		switch {
		case s[mid] == key:
			fmt.Printf("The position of %d is %d\n", key, mid)
			return mid
		case s[mid] < key:
			low = mid + 1
			fmt.Printf("New low = %d\n", low)
		case s[mid] > key:
			high = mid - 1
			fmt.Printf("New high = %d\n", high)
		}
		/*
			if s[mid] == key {
				fmt.Printf("The position of %d is %d\n", key, mid)
				return mid
			} else if s[mid] < key {
				low = mid + 1
				fmt.Printf("New low = %d\n", low)
			} else if {
				high = mid - 1
				fmt.Printf("New high = %d\n", high)
			}
		*/
	}
	return -1
}
