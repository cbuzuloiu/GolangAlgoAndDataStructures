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
	fmt.Printf("The slice will be s = %v\n", os.Args[1:len(os.Args)-1])
	fmt.Printf("The slice type is: %T\n", os.Args[1:])
	fmt.Printf("The key is: %v\n", os.Args[len(os.Args)-1])
	fmt.Printf("Because s is a slice of strings we need to convert it to a slice of int\n")

	s := make([]int, 0, len(os.Args)-2)

	key, err := strconv.Atoi(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal("Failed to convert string to integer ", err)
	}

	fmt.Printf("key = %d\n", key)

	for _, arg := range os.Args[1 : len(os.Args)-1] {
		intArg, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatal("Failed to convert string to integer ", err)
		}

		s = append(s, intArg)
	}

	fmt.Printf("s = %v\n", s)

	targetValue := binarySearch(s, key)

	if targetValue == -1 {
		fmt.Println("The target value is not in the array")
	} else {
		fmt.Printf("The target value is at position: %d\n", targetValue)
	}
}

func binarySearch(s []int, key int) int {
	low := 0
	high := len(s) - 1

	for low <= high {
		// avoid overflow
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
