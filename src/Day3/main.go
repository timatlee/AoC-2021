package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	commands := readfile("testinput.txt")
	testdata_power := find_power_consumption(commands)
	println("Part 1 test data (should be 198): ", testdata_power)
	println("Part 2 test data (should be 230): ", find_life_support_rating(commands))

	realCommands := readfile("day3input.txt")
	realdata_power := find_power_consumption(realCommands)
	println("Part 1 actual data: ", realdata_power)
	println("Part 2 actual data: ", find_life_support_rating(realCommands))

}

func find_string_with_value_in_index(commands []string, index int, value byte) []string {
	var s []string
	for i := 0; i < len(commands); i++ {
		command := commands[i]
		if command[index] == value {
			s = append(s, command)
		}
	}
	return s
}

func count_values_in_index(commands []string, index int) (int, int) {
	var zeros, ones int = 0, 0
	// Iterate over the list of strings
	for i := 0; i < len(commands); i++ {
		// Test each string's index's value
		if commands[i][index] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	return zeros, ones
}

func find_life_support_rating(commands []string) int {
	// make a copy of the original commands
	o := append([]string{}, commands...)

	// Loop over the list of MSB bits in Oxygen.
	for i := 0; i < len(commands[0]); i++ {
		// Find how many 0s and 1s there are in the list of commands
		number0, number1 := count_values_in_index(o, i)
		//fmt.Print("index", i, number0, number1)

		if number0 > number1 {
			o = find_string_with_value_in_index(o, i, byte('0'))
		} else if number1 > number0 {
			o = find_string_with_value_in_index(o, i, byte('1'))
		} else if number1 == number0 {
			o = find_string_with_value_in_index(o, i, byte('1'))
		}

		//fmt.Println(o)
		if len(o) < 2 {
			break
		}
	}

	c := append([]string{}, commands...)
	// Loop over the list of :SB bits in CO2.
	for i := 0; i < len(commands[0]); i++ {
		// Find how many 0s and 1s there are in the list of commands
		number0, number1 := count_values_in_index(c, i)
		//fmt.Print("index", i, number0, number1)

		if number0 > number1 {
			c = find_string_with_value_in_index(c, i, byte('1'))
		} else if number1 > number0 {
			c = find_string_with_value_in_index(c, i, byte('0'))
		} else if number1 == number0 {
			c = find_string_with_value_in_index(c, i, byte('0'))
		}

		//fmt.Println(c)
		if len(c) < 2 {
			break
		}
	}

	oxygen_value, _ := strconv.ParseInt(o[0], 2, 64)
	cotwo_value, _ := strconv.ParseInt(c[0], 2, 64)

	return int(oxygen_value) * int(cotwo_value)
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
