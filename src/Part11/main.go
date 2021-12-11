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
	filecontent := readfile("test1input.txt")
	//filecontent := readfile("day10input.txt")

	octopus := parseInput(filecontent)
	PrintOctopus(octopus)
	increaseEnergy(octopus)

}

func PrintOctopus(o [][]Octopus) {
	for y := 0; y < len(o); y++ {
		fmt.Println(o[y])
	}

}

func increaseEnergy(o [][]Octopus) {
	for y := 0; y < len(o); y++ {
		for x := 0; x < len(o[y]); x++ {
			o[y][x].Energy += 1
			o[y][x].HasFlashed = false
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
			o[x][y].Energy = digit
			o[x][y].HasFlashed = false
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
