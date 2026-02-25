package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run file_name.go [test_data.txt]")
	}

	testDataFiles := os.Args[1:]

	fmt.Printf("Data files: %v\n", testDataFiles)

	for _, file := range testDataFiles {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v\n", err)
			continue
		}

		scanner := bufio.NewScanner(f)

		// Discard the header
		// We scan once to advance the cursor past the first line
		if !scanner.Scan() {
			// If this fails the input is empty
			continue
		}

		dataResult := make(chan *DataStruct)
		nrOfWorkeres := 0
		var receivedDataSlice []*DataStruct

		for scanner.Scan() {
			// stringToIntSlice(scanner.Text())
			go workerGorutine(scanner.Text(), file, dataResult)
			nrOfWorkeres++
			/*
				fmt.Printf("%T", dataResult)
				receivedData := <-dataResult
				fmt.Println("---- **** ----")
				fmt.Printf("For file name: %v\nNr Crt: %d\nThe data Slice is: %v\nThe target Key is: %d\nThe index of the target is: %d\n", receivedData.FileName, receivedData.NrCrt, receivedData.DataSlice, receivedData.Key, receivedData.IndexOfTarget)
				fmt.Println("---- **** ----")
			*/
		}

		for i := 0; i < nrOfWorkeres; i++ {
			// receivedData := <-dataResult
			receivedDataSlice = append(receivedDataSlice, <-dataResult)
		}

		fmt.Println(receivedDataSlice)

		for _, result := range receivedDataSlice {
			fmt.Println("---- **** ----")
			fmt.Printf("For file name: %v\nNr Crt: %d\nThe data Slice is: %v\nThe target Key is: %d\nThe index of the target is: %d\n", result.FileName, result.NrCrt, result.DataSlice, result.Key, result.IndexOfTarget)
			fmt.Println("---- **** ----")
		}
	}
	fmt.Println(time.Since(start).Seconds())
}

type DataStruct struct {
	FileName      string `json:"fileName"`
	Key           int    `json:"key"`
	DataSlice     []int  `json:"dataSlice"`
	NrCrt         int    `json:"nrCrt"`
	IndexOfTarget int    `json:"IndexOfTarget"`
}

func workerGorutine(data string, fileName string, ch chan<- *DataStruct) {
	s := stringToIntSlice(data)

	key := s[len(s)-1]
	dataS := s[1 : len(s)-2]
	nrCrt := s[0]

	/*
		fmt.Printf("File name: %s\n", fileName)
		fmt.Printf("Key: %d\n", key)
		fmt.Printf("Data slice : %v\n", dataS)
		fmt.Printf("Nr Crt: %d\n", nrCrt)
	*/

	dataStruct := &DataStruct{
		FileName:  fileName,
		Key:       key,
		DataSlice: dataS,
		NrCrt:     nrCrt,
	}

	dataStruct.IndexOfTarget = binarySearch(dataS, key)

	/*
		fmt.Println("The Data Struct is: ")
		pData, err := json.MarshalIndent(dataStruct, "", "    ")
		if err != nil {
			fmt.Println("Error marshaling:", err)
		}

		fmt.Println(string(pData))
	*/

	ch <- dataStruct
}

func stringToIntSlice(data string) []int {
	// Split into a slice of string split by whitespace (spaces, tabs, newlines)
	strSlice := strings.Fields(data)

	// Pre allocate the integer slice
	intSlice := make([]int, 0, len(strSlice))

	// Convert the strSlice to intSlice
	for _, elem := range strSlice {
		num, err := strconv.Atoi(elem)
		if err != nil {
			log.Fatal(err)
		}
		intSlice = append(intSlice, num)
	}

	return intSlice
}

func binarySearch(s []int, key int) int {
	low := 0
	high := len(s) - 1

	for low <= high {
		mid := low + (high-low)/2

		switch {
		case s[mid] == key:
			return mid
		case s[mid] < key:
			low = mid + 1
		default:
			high = mid - 1
		}
	}

	return -1
}
