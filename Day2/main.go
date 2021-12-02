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
	commands := readfile("testinput.txt")
	realCommands := readfile("day2input.txt")

	println("Part 1 test data (should be 150): ", find_total_position(commands))
	println("Part 1 actual data: ", find_total_position(realCommands))

	// Part 2
	/*
		println("Part 2 test data: ", find_deeper_count_sliding_window(numbers, 3))
		println("test?", test(numbers))
		println("Part 2 actual data: ", find_deeper_count_sliding_window(realNumbers, 3))
		println("Part 2 actual data: ", test(realNumbers))
	*/
}

func find_total_position(commands []string) int {
	var pos int = 0
	var depth int = 0

	for _, rawcommand := range commands {
		fmt.Println(rawcommand)
		split := strings.Split(rawcommand, " ")
		command := split[0]
		distance, _ := strconv.Atoi(split[1])

		switch command {
		case "up":
			depth -= distance
		case "down":
			depth += distance
		case "forward":
			pos += distance
		}

	}
	return depth * pos
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
