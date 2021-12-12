package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Octopus struct {
	Energy     int
	HasFlashed bool
}

func main() {
	//filecontent := readfile("test1input.txt")
	//filecontent := readfile("testinput.txt")
	filecontent := readfile("day11input.txt")

	octopus := parseInput(filecontent)
	// Inital state
	PrintOctopus(octopus)
	var totalFlashes int

	for i := 0; i < 100; i++ {
		totalFlashes += stepOne(octopus)
	}
	fmt.Println("Part 1:", totalFlashes)

	// Part 2 ..
	octopus2 := parseInput(filecontent)
	var stepCounter int = 0
	for {
		stepCounter += 1
		stepOne(octopus2)
		if partTwo(octopus2) {
			fmt.Println(stepCounter)
			break
		}
	}

}

func partTwo(o [][]Octopus) bool {
	var allZeros bool = true
	for y := 0; y < len(o); y++ {
		for x := 0; x < len(o[y]); x++ {
			if o[y][x].Energy != 0 {
				allZeros = false
				return allZeros
			}
		}
	}

	return allZeros
}

func stepOne(o [][]Octopus) int {
	// FIrst phase: Increase the energy level on all O's.
	increaseEnergy(o)

	newFlashes := 1
	totalFlashes := 0
	for newFlashes > 0 {
		newFlashes = 0

		for y := 0; y < len(o); y++ {
			for x := 0; x < len(o[y]); x++ {
				if o[y][x].Energy > 9 && !o[y][x].HasFlashed {
					newFlashes += 1
					o[y][x].HasFlashed = true
					increaseAdjacents(o, x, y)
				}
			}
		}
		totalFlashes += newFlashes
	}
	// Cleanup.
	for y := 0; y < len(o); y++ {
		for x := 0; x < len(o[y]); x++ {
			if o[y][x].Energy > 9 {
				o[y][x].Energy = 0
				o[y][x].HasFlashed = false
			}
		}
	}
	return totalFlashes
}

func increaseAdjacents(o [][]Octopus, x int, y int) int {
	var newFlashes int = 0
	for a := -1; a <= 1; a++ { // Y offset
		for b := -1; b <= 1; b++ { // X offset
			if !(a == 0 && b == 0) { // Check that we're not in the middle.
				// Check that we're within Y bounds.
				if (y+a >= 0) && (y+a < len(o)) {
					// Check that we're within the X bounds
					if (x+b >= 0) && (x+b < len(o[y])) {
						// Increase this offset'd point. If it's > 9, indicate a flash.
						o[y+a][x+b].Energy++
					}
				}
			}
		}
	}
	return newFlashes
}

func PrintOctopus(o [][]Octopus) {
	for y := 0; y < len(o); y++ {
		fmt.Println(o[y])
	}
	fmt.Println()

}

func increaseEnergy(o [][]Octopus) {
	for y := 0; y < len(o); y++ {
		for x := 0; x < len(o[y]); x++ {
			o[y][x].Energy += 1
		}
	}
}

func parseInput(f []string) [][]Octopus {
	var o [][]Octopus = make([][]Octopus, len(f))
	for i := range o {
		o[i] = make([]Octopus, len(f[i]))
	}
	for y := 0; y < len(f); y++ {
		for x := 0; x < len(f[y]); x++ {
			digit, _ := strconv.Atoi(string(f[y][x]))
			o[y][x].Energy = digit
			o[y][x].HasFlashed = false
		}
	}

	return o
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
		text := scanner.Text()
		if len(text) > 0 {
			commands = append(commands, text)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return commands
}
