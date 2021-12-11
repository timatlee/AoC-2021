package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	//filecontent := readfile("testinput.txt")
	filecontent := readfile("day10input.txt")

	var incompleteLines []string

	var badrunes = make(map[rune]int)
	var scores int
	for _, v := range filecontent {
		badrune := parseSyntax(v)
		if badrune != '0' {
			badrunes[badrune]++
		} else {
			incompleteLines = append(incompleteLines, v)
		}
	}
	for i, v := range badrunes {
		switch i {
		case ')':
			scores += (v * 3)
		case ']':
			scores += (v * 57)
		case '}':
			scores += (v * 1197)
		case '>':
			scores += (v * 25137)
		}
	}
	fmt.Printf("Part 1 scores: %v\n", scores)

	// Part 2
	var pt2scores []int
	for _, v := range incompleteLines {
		closers := completeLines(v)
		score := 0

		for _, v := range closers {
			score *= 5

			switch v {
			case ')':
				score += 1
			case ']':
				score += 2
			case '}':
				score += 3
			case '>':
				score += 4
			}
		}

		fmt.Println(string(closers), score)
		pt2scores = append(pt2scores, score)
	}
	sort.Ints(pt2scores)
	middle := pt2scores[len(pt2scores)/2]

	fmt.Println("Part 2 middle:", middle)
}

func completeLines(line string) []rune {
	var commands []rune
	// Iterate the characters in the string
	for _, v := range line {
		if len(commands) < 1 {
			commands = append(commands, v)
		} else {
			lastChar := commands[len(commands)-1]
			if v == rune('(') || v == rune('[') || v == rune('{') || v == rune('<') {
				// If current character is an open rune, we can just append it.
				commands = append(commands, v)
			} else {
				// Otherwise, it's a closing character and we need to check it.
				closeChar := getMatchingClose(lastChar)

				if closeChar == v {
					// Pop off the last element.
					commands = commands[:len(commands)-1]
					// Commands now contains unclosed runes.
				}
			}
		}
	}
	// Commands now contains the remaining open runes.
	// Work backwards from the end of commands, and append what we need.
	var closers []rune
	for i := len(commands) - 1; i >= 0; i-- {
		closers = append(closers, getMatchingClose(commands[i]))
	}

	fmt.Println("Command:", string(commands), "Closed by:", string(closers))
	return closers
}

func parseSyntax(line string) rune {
	var commands []rune
	var badrune rune = '0'
	for _, v := range line {
		// First index gets appended automatically because we can't check the previous entry
		if len(commands) < 1 {
			commands = append(commands, v)
		} else {
			lastChar := commands[len(commands)-1]
			if v == rune('(') || v == rune('[') || v == rune('{') || v == rune('<') {
				// If current character is an open rune, we can just append it.
				commands = append(commands, v)
			} else {
				// Otherwise, it's a closing character and we need to check it.
				closeChar := getMatchingClose(lastChar)

				if closeChar == v {
					commands = commands[:len(commands)-1]
				} else {
					badrune = v
					break
				}
			}
		}
	}
	return badrune
	//	fmt.Printf("commands: %v\n", string(commands))
}

func getMatchingClose(s rune) rune {
	switch string(s) {
	case "(":
		return ')'
	case "[":
		return ']'
	case "{":
		return '}'
	case "<":
		return '>'
	}

	return 'x'
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
