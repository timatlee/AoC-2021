package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	commands := readfile("testinput.txt")
	realCommands := readfile("day3input.txt")

	println("Part 1 test data (should be 198): ", find_power_consumption(commands))
	println("Part 1 actual data: ", find_power_consumption(realCommands))

	// Part 2
	/*
		println("Part 2 test data: ", find_deeper_count_sliding_window(numbers, 3))
		println("test?", test(numbers))
		println("Part 2 actual data: ", find_deeper_count_sliding_window(realNumbers, 3))
		println("Part 2 actual data: ", test(realNumbers))
	*/
}

func find_power_consumption(commands []string) int {
	position_ones := make(map[int]int)
	position_zeros := make(map[int]int)

	for _, value := range commands {
		for command_index, command_value := range value {
			// We seem to be getting the ASCII character value.  Offset by 48 because I'm lazy.
			binary_number := command_value - 48

			if binary_number == 1 {
				position_ones[command_index] += 1
			} else {
				position_zeros[command_index] += 1
			}
		}
	}

	// fmt.Println(position_ones)
	// fmt.Println(position_zeros)

	// find which entry, one / zero, is more significant.
	var gamma, epsilon string

	for index := 0; index < len(position_ones); index++ {
		if position_ones[index] > position_zeros[index] {
			// fmt.Println("Most common in position ", index, ": 1")
			gamma += "1"
			epsilon += "0"
		} else {
			// fmt.Println("Most common in position ", index, ": 0")
			gamma += "0"
			epsilon += "1"
		}
	}
	gamma_value, _ := strconv.ParseInt(gamma, 2, 64)
	epsilon_value, _ := strconv.ParseInt(epsilon, 2, 64)
	// fmt.Println("Gamma", gamma, "Epsilon", epsilon)
	// fmt.Println("Gamma", gamma_value, "Epsilon", epsilon_value)
	return int(gamma_value) * int(epsilon_value)
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
