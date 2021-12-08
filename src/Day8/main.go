package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	//filecontent := readfile("testinput.txt")
	filecontent := readfile("day8input.txt")

	var digitCounts = make([]int, 10)

	for i, _ := range filecontent {
		_, digits := splitSignalsAndDigits(filecontent[i])
		for _, v := range digits {
			digitLengs := len(v)
			actualDigit := mapSegcountToDigit(digitLengs)
			digitCounts[actualDigit]++
		}
	}

	fmt.Printf("digitCounts: %v\n", digitCounts)
	// element 0 is our "unknown" number, so we're not going to usm it.
	fmt.Println("Part 1:", sumArray(digitCounts[1:]))

	// Part 2
	var numbersum int
	for i, _ := range filecontent {
		input, output := splitSignalsAndDigits(filecontent[i])
		digitIndex := determine_letter_to_number(input)

		var sequences []string
		var numbers []int
		for _, v := range output {
			sequences = append(sequences, v)
			numbers = append(numbers, indexOfArray(digitIndex, v))
		}
		number := sliceToInt(numbers)
		numbersum += number
	}
	fmt.Println(numbersum)
}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

func indexOfArray(s []string, item string) int {
	for i := range s {
		if s[i] == item {
			return i
		}
	}

	return -1
}

// https://www.reddit.com/r/adventofcode/comments/rbvpui/2021_day_8_part_2_my_logic_on_paper_i_used_python/
//
// What the fuck is going on here.
// Take our 10 inputs, and return an array of 10 integers that map each index to the letter code.
func determine_letter_to_number(input []string) []string {
	var letters = make([]string, 10)

	// First identify 1, 4, 7 and 8. We can do this by the length, used in part 1.
	for _, v := range input {
		digit := mapSegcountToDigit(len(v))
		// Anything we don't know is coming in as a 0. We might have to fix this later.
		if digit != 0 {
			letters[digit] = v
		}
	}

	// Find the differences between 1 and 4. We can use this for a bunch of other numbers.
	fourDiff := difference(letters[4], letters[1])

	for _, v := range input {
		// Thinking about the 5 segment options now,
		if len(v) == 5 {
			// To identify 3, if an input contains both elements of 1
			if find_match_in_input(letters[1], v) {
				letters[3] = v
			}

			// To identify 5, if an input contains "fourdiff"
			if find_match_in_input(fourDiff, v) {
				letters[5] = v
			}

			// To identify a 2, if an input does not match the parameters for 3 and 5.
			if !find_match_in_input(letters[1], v) && !find_match_in_input(fourDiff, v) {
				letters[2] = v
			}
		}

		if len(v) == 6 {
			// To identify 9, all segments in 9 and 4 are the same.
			if find_match_in_input(letters[4], v) {
				letters[9] = v
			}

			// To identify 6, it should match fourDiff, but NOT 9.
			if find_match_in_input(fourDiff, v) && !find_match_in_input(letters[4], v) {
				letters[6] = v
			}

			// For 0, it's not 9 or 6.
			if !(find_match_in_input(fourDiff, v) && !find_match_in_input(letters[4], v)) && !find_match_in_input(letters[4], v) {
				letters[0] = v
			}
		}
	}

	return letters
}

func find_match_in_input(source, input string) bool {
	onemap := make(map[rune]int, len(source))

	for _, x := range source {
		onemap[x] = 0
	}

	for _, v := range input {
		_, ok := onemap[v]
		if ok {
			onemap[v]++
		}
	}

	for _, v := range onemap {
		if v == 0 {
			return false
		}
	}

	return true
}

// difference returns the elements in `a` that aren't in `b`.
func difference(a, b string) string {
	mb := make(map[rune]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []rune
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return string(diff)
}

func mapSegcountToDigit(s int) int {
	switch s {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	}

	return 0
}

func sumArray(fish []int) int {
	var sum int
	for _, v := range fish {
		sum += v
	}
	return sum
}

func splitSignalsAndDigits(l string) ([]string, []string) {
	var signals, outputs []string
	// Split the input line by |
	datas := strings.Split(l, "|")

	// Then split each group by space
	sigs := strings.Split(datas[0], " ")
	for i := range sigs {
		if len(sigs[i]) > 0 {
			s := []rune(sigs[i])
			sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
			signals = append(signals, string(s))
		}
	}
	outs := strings.Split(datas[1], " ")
	for i := range outs {
		if len(outs[i]) > 0 {
			s := []rune(outs[i])
			sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
			outputs = append(outputs, string(s))
		}
	}

	return signals, outputs
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
