package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := readfile()
	fmt.Println(numbers)
}

func readfile() []int {
	file, err := os.Open("day1/testinput.txt")
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
