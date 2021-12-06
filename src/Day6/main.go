package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//filecontent := readfile("testinput.txt")
	filecontent := readfile("day6input.txt")
	var days int = 256

	fishStarts := strings.Split(filecontent[0], ",")

	var fishInternalTimer [9]int
	// Set initial age
	for _, v := range fishStarts {
		vInt, _ := strconv.Atoi(v)
		fishInternalTimer[vInt] += 1
	}

	fmt.Printf("fishInternalTimer: %v\n", fishInternalTimer)

	for i := 0; i < days; i++ {
		growFish(fishInternalTimer[:])
		fmt.Printf("fishInternalTimer: %v\n", fishInternalTimer)
	}

	fmt.Printf("sumArray(fishInternalTimer[:]): %v\n", sumArray(fishInternalTimer[:]))
}

func sumArray(fish []int) int {
	var sum int
	for _, v := range fish {
		sum += v
	}
	return sum
}

func growFish(fish []int) []int {
	// Deal with fish with timer of 0.
	var newFish6, newFish8 int

	if fish[0] > 0 {
		// The number of fish at 0 now becomes timer 6.
		fishAt0 := fish[0]
		fish[0] = 0
		newFish6 = fishAt0

		// Create new fish with internal timer of 8.
		newFish8 = fishAt0
	}

	// Counting down the non-0 timer fihs.
	for i := 1; i < len(fish); i++ {
		if fish[i] > 0 {
			fish[i-1] += fish[i]
			fish[i] = 0
		}
	}

	// Add the growth at 6 and 8
	fish[6] += newFish6
	fish[8] += newFish8

	return fish

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
