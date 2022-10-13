package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := readfile("testinput.txt")
	realNumbers := readfile("day1input.txt")

	println("Part 1 test data: ", find_deeper_count(numbers))
	println("Part 1 actual data: ", find_deeper_count(realNumbers))

	// Part 2
	println("Part 2 test data: ", find_deeper_count_sliding_window(numbers, 3))
	println("test?", test(numbers))
	println("Part 2 actual data: ", find_deeper_count_sliding_window(realNumbers, 3))
	println("Part 2 actual data: ", test(realNumbers))
}

func test(numbers []int) int {
	var sliceSize = 3
	var indexStart int = 0
	var indexEnd int = indexStart + sliceSize
	var increaseCounter = 0

	for range numbers {
		if numbers[indexEnd] > numbers[indexStart] {
			increaseCounter++
		}

		indexStart++
		indexEnd++

		if indexEnd >= len(numbers) {
			break
		}
	}

	return increaseCounter
}

func find_deeper_count_sliding_window(numbers []int, sliceSize int) int {
	var indexStart int = 0
	var indexEnd int = indexStart + sliceSize
	var lastSliceSum int = 0
	var thisSliceSum int
	var increaseCounter = 0

	for range numbers {
		lastSliceSum = thisSliceSum
		thisSliceSum = addArray(numbers[indexStart:indexEnd])
		if lastSliceSum == 0 {
			lastSliceSum = thisSliceSum
		}

		if thisSliceSum > lastSliceSum {
			increaseCounter++
		}

		indexStart++
		indexEnd++

		if indexEnd > len(numbers) {
			break
		}
	}

	return increaseCounter
}

func addArray(numbs []int) int {
	result := 0
	for _, numb := range numbs {
		result += numb
	}
	return result
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
