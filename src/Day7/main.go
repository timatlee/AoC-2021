package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
)

func main() {
	//filecontent := readfile("testinput.txt")
	filecontent := readfile("day7input.txt")
	inputStrings := strings.Split(filecontent[0], ",")

	var input []int
	for _, v := range inputStrings {
		number, _ := strconv.Atoi(v)
		input = append(input, number)
	}

	data := stats.LoadRawData(input)
	//mean, _ := stats.Mean(data)
	median, _ := stats.Median(data)
	//mode, _ := stats.Mode(data)

	//fmt.Println(mean)   // Average
	//fmt.Println(median) // Middle number in the sorted set
	//fmt.Println(mode)   // Most occuring number

	// Part 1
	var fuelSum int
	for _, v := range input {
		fuelDiff := math.Abs(float64(v) - median)
		fuelSum += int(fuelDiff)
	}
	fmt.Printf("Part 1 fuelSum: %v\n", fuelSum)

	// Part 2. Brute force I think,.
	iters := 1000

	var itersList = make([]int, iters)
	for i := 0; i < iters; i++ {
		var fuelCost int = 0
		for _, v := range input {
			tempFuelCost := addBetweenRanges(v, i)
			fuelCost += tempFuelCost
		}
		itersList[i] = fuelCost
	}
	//fmt.Printf("itersList: %v\n", itersList)
	minValue, minValueIndex := findMinValueInArray(itersList)
	fmt.Println("Part 2: Value,Index:", minValue, minValueIndex)
}

func findMinValueInArray(s []int) (int, int) {
	var minValue int = s[0]
	var minIndex int = 0
	for i, v := range s {
		if s[i] < minValue {
			minValue = v
			minIndex = i
		}
	}

	return minValue, minIndex
}

func addBetweenRanges(a int, b int) int {
	var tempSum int
	for i := 0; i <= int(math.Abs(float64(b-a))); i++ {
		tempSum += i
	}
	return tempSum
}

func readfile(filename string) []string {
	dir, _ := os.Getwd()

	file, err := os.Open(dir + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var commands []string

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return commands
}
