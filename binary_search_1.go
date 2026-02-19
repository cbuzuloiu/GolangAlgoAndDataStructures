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
		// see below commnet about mid
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

		// --- if else ---
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

/*
Advanced Consideration: The Integer Overflow Bug

While your code is perfectly fine for normal use cases, there is a famous historical bug associated
with the line mid := (high + low) / 2.

If you are working with a massively large array (specifically, an array size close to the maximum
limit of an integer), adding high and low together can cause an integer overflow before the division happens.
This results in a negative number for mid, which will cause an index out-of-bounds panic.

The Fix:
To make your binary search mathematically indestructible, you can calculate the midpoint by taking the distance
between high and low, dividing that by two, and adding it to low:
*/
