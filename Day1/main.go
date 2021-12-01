package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := readfile("testinput.txt")

	var testIncreases = find_deeper_count(numbers)
	println(testIncreases)

	var testSliceIncreases = find_deeper_count_sliding_window(numbers, 3)
	println(testSliceIncreases)

	realNumbers := readfile("day1input.txt")
	var increases = find_deeper_count(realNumbers)
	println(increases)

}

func find_deeper_count_sliding_window(numbers []int, sliceSize int) int {
	var thisSlice int = 0
	var lastSlice int = 0
	var increases int = 0

	for index, number := range numbers {
		//	fmt.Println(index, number)
		if index >= 0 && (index%sliceSize == 0) {
			if lastSlice > thisSlice {
				increases++
			}
			fmt.Println(thisSlice)
			lastSlice = thisSlice
			thisSlice = 0

		}
		thisSlice += number
	}

	return increases

}

func find_deeper_count(numbers []int) int {
	var lastNumber int = 0
	var increaseCounter int = 0

	for _, number := range numbers {
		//	fmt.Println(index, number)
		if lastNumber != 0 && number > lastNumber {
			increaseCounter++
		}
		lastNumber = number
	}

	return increaseCounter
}

func readfile(filename string) []int {
	dir, _ := os.Getwd()

	file, err := os.Open(dir + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numbers []int

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		intVar, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, intVar)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numbers
}
