package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := readfile("testinput.txt")

	var testIncreases = find_deeper_count(numbers)
	println(testIncreases)

	realNumbers := readfile("day1input.txt")
	var increases = find_deeper_count(realNumbers)
	println(increases)

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
