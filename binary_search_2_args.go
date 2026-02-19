package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("\t-----\tBINARY SEARCH\t-----")
	fmt.Printf("The slice will be s = %v\n", os.Args[1:len(os.Args)-1])
	fmt.Printf("The slice type is: %T\n", os.Args[1:])
	fmt.Printf("The key is: %v\n", os.Args[len(os.Args)-1])
	fmt.Printf("Because s is a slice of strings we need to convert it to a slice of int\n")

	var s []int

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
}
