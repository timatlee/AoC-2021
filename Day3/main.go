package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	commands := readfile("testinput.txt")

	testdata_power, testdata_oxygen, testdata_cotwo := find_power_consumption(commands)
	println("Part 1 test data (should be 198): ", testdata_power)
	println("Part 2 test data (should be 230): ", find_life_support_rating(commands, testdata_oxygen, testdata_cotwo))

	/*
		realCommands := readfile("day3input.txt")
		realdata_power, realdata_oxygen, realdata_cotwo := find_power_consumption(realCommands)
		println("Part 1 actual data: ", realdata_power)
		println("Part 2 actual data: ", find_life_support_rating(realCommands, realdata_oxygen, realdata_cotwo))
	*/
}

func find_life_support_rating(commands []string, oxygen string, cotwo string) int {
	// Commands related to oxygen system.
	oxygen_commands := append([]string{}, commands...)
	cotwo_commands := append([]string{}, commands...)

	// Loop over the list of MSB bits in Oxygen.
	for i := 0; i < len(oxygen); i++ {
		// MSB of all the bits in this position.
		msbbit := oxygen[i]

		// Commands that we're keeping from this loop.
		var kept_commands []string

		// fmt.Println("Bitness to keep:", msbbit-48, "in position", i)

		// Check if there are more than two commands left.  If so..
		if len(oxygen_commands) > 2 {
			// Now loop over all our commands
			for j := 0; j < len(oxygen_commands); j++ {
				// Test if a command, bit position i, matches.
				if oxygen_commands[j][i] == msbbit {
					kept_commands = append(kept_commands, oxygen_commands[j])
				}
			}
		} else if len(oxygen_commands) == 2 {
			// If not, we need to find which command is "bigger".
			first, _ := strconv.ParseInt(oxygen_commands[0], 2, 64)
			second, _ := strconv.ParseInt(oxygen_commands[1], 2, 64)
			if first > second {
				kept_commands = append(kept_commands, oxygen_commands[0])
			} else {
				kept_commands = append(kept_commands, oxygen_commands[1])
			}
		} else {
			break
		}
		oxygen_commands = kept_commands
		// fmt.Println("Oxygen MSB index", i, "MSB Bit", msbbit-48, "remaining commands", oxygen_commands)
	}

	// Loop over the list of MSB bits in CO2.
	for i := 0; i < len(cotwo); i++ {
		// MSB of all the bits in this position.
		msbbit := cotwo[i]

		// Commands that we're keeping from this loop.
		var kept_commands []string

		// fmt.Println("Bitness to keep:", msbbit-48, "in position", i)

		// Check if there are more than two commands left.  If so..
		if len(cotwo_commands) > 2 {
			// Now loop over all our commands
			for j := 0; j < len(cotwo_commands); j++ {
				// Test if a command, bit position i, matches.
				if cotwo_commands[j][i] == msbbit {
					kept_commands = append(kept_commands, cotwo_commands[j])
				}
			}
		} else if len(cotwo_commands) == 2 {
			// If not, we need to find which command is "bigger".
			first, _ := strconv.ParseInt(cotwo_commands[0], 2, 64)
			second, _ := strconv.ParseInt(cotwo_commands[1], 2, 64)
			if first < second {
				kept_commands = append(kept_commands, cotwo_commands[0])
			} else {
				kept_commands = append(kept_commands, cotwo_commands[1])
			}
		} else {
			break
		}
		cotwo_commands = kept_commands
		//fmt.Println("CO2 MSB index", i, "MSB Bit", msbbit-48, "remaining commands", cotwo_commands)
	}

	// oxygen_commands now contains the binary value of our oxygen.
	// cotwo_commands contains the binary value of our co2.

	oxygen_value, _ := strconv.ParseInt(oxygen_commands[0], 2, 64)
	cotwo_value, _ := strconv.ParseInt(cotwo_commands[0], 2, 64)

	fmt.Println(oxygen_commands[0], cotwo_commands[0], oxygen_value, cotwo_value)

	return int(oxygen_value) * int(cotwo_value)
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)

	/*
		copy(slice[s:], slice[s+1:])
		slice[len(slice)-1] = ""
		slice = slice[:len(slice)-1]
	*/

	/*
		var newslice []string
		for i := 0; i < len(slice); i++ {
			if i != s {
				newslice = append(newslice, slice[i])
			}
		}
		return newslice
	*/
}

func find_power_consumption(commands []string) (int, string, string) {
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
	return int(gamma_value) * int(epsilon_value), gamma, epsilon
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
