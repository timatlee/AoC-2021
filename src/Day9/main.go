package main

import (
	"Day9/vertex"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//filecontent := readfile("testinput.txt")
	filecontent := readfile("day9input.txt")

	// Figure out the sea floor
	seafloor := parseContentTo2dArray(filecontent)

	// Find low spots
	lowspots := findLowSpots(seafloor)

	// Find the value of the low spots, and figure out risk.
	var sumrisk, risk int
	for _, v := range lowspots {
		risk = seafloor[v.Y][v.X]
		// checkPoint(seafloor, v)
		sumrisk += risk + 1

	}
	fmt.Printf("risk: %v\n", sumrisk)
}

func checkPoint(s [][]int, v vertex.Vertex) {
	fmt.Println("Current:", s[v.Y][v.X], "at x,y", v.X, v.Y)
	if v.Y > 0 {
		fmt.Println("Up:", s[v.Y-1][v.X], "at x,y", v.X, v.Y-1)
	}
	if v.Y < len(s)-1 {
		fmt.Println("Down:", s[v.Y+1][v.X], "at x,y", v.X, v.Y+1)
	}
	if v.X > 0 {
		fmt.Println("Left:", s[v.Y][v.X-1], "at x,y", v.X-1, v.Y)
	}
	if v.X < len(s[0])-1 {
		fmt.Println("Right:", s[v.Y][v.X+1], "at x,y", v.X+1, v.Y)
	}
	fmt.Println()
}

func findLowSpots(seafloor [][]int) []vertex.Vertex {
	var lowspots []vertex.Vertex

	for y := range seafloor {
		for x := range seafloor[y] {
			isLowest := true
			c := seafloor[y][x]

			// Check left
			if x > 0 {
				v := seafloor[y][x-1]
				if c >= v {
					isLowest = false
				}
			}
			// Check right
			if x < len(seafloor[y])-1 && isLowest {
				v := seafloor[y][x+1]
				if c >= v {
					isLowest = false
				}
			}

			// Check up
			if y > 0 && isLowest {
				v := seafloor[y-1][x]
				if c >= v {
					isLowest = false
				}
			}

			// Check down
			if y < len(seafloor)-1 && isLowest {
				v := seafloor[y+1][x]
				if c >= v {
					isLowest = false
				}
			}

			if isLowest {
				lowspots = append(lowspots, vertex.NewFromInt(x, y))
			}
		}
	}

	return lowspots
}

func parseContentTo2dArray(content []string) [][]int {
	var seafloor [][]int
	for i := range content {
		var row = make([]int, len(content[i]))
		for j := range content[i] {
			value, _ := strconv.Atoi(string(content[i][j]))
			row[j] = value
		}
		seafloor = append(seafloor, row)
	}

	return seafloor
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
