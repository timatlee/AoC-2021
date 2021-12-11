package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//filecontent := readfile("testinput.txt")
	filecontent := readfile("day10input.txt")

	var badrunes = make(map[rune]int)
	var scores int
	for _, v := range filecontent {
		badrune := parseSyntax(v)
		if badrune != '0' {
			badrunes[badrune]++
		}
	}
	for i, v := range badrunes {
		fmt.Println("Char:", string(i), "value:", v)
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
	fmt.Printf("scores: %v\n", scores)
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
