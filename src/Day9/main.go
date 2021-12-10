package main

import (
	"Day9/vertex"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	// Part 2.  First..  I'm going to give the perimeter of the seafloor a height of 9.
	// This makes basins end at the perimeter, effectively.
	// Did this in parseContentto2darray() function.

	// Next, iterate over our known lowspots (basins)
	var basinSizes []int
	for _, v := range lowspots {
		// Count low spots in the basin.
		foo := countLowSpots(seafloor, v.X, v.Y, 0)
		basinSizes = append(basinSizes, foo)
	}
	sort.Ints(basinSizes)

	parts := basinSizes[len(basinSizes)-3:]
	fmt.Printf("parts: %v\n", parts)

	fmt.Println(parts[0] * parts[1] * parts[2])
}

func countLowSpots(m [][]int, x int, y int, i int) int {
	// Check if the value at this coordinate is 9. If so, we're done here.
	v := m[y][x]
	if v != 9 {
		i += 1
		// Recursion!
		// Check the coordinates around this point.
		m[y][x] = 9
		i = countLowSpots(m, x, y+1, i)
		i = countLowSpots(m, x, y-1, i)
		i = countLowSpots(m, x-1, y, i)
		i = countLowSpots(m, x+1, y, i)
	}

	return i

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

	for y := 1; y < len(seafloor)-1; y++ {
		for x := 1; x < len(seafloor[y])-1; x++ {
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
	var zeroRow = make([]int, len(content[1])+2)
	for i := range zeroRow {
		zeroRow[i] = 9
	}

	zeroRowCopy := make([]int, len(zeroRow))
	copy(zeroRowCopy, zeroRow)

	seafloor = append(seafloor, zeroRow)
	for i := range content {
		var row = make([]int, len(content[i])+2)

		row[0] = 9
		for j := range content[i] {
			value, _ := strconv.Atoi(string(content[i][j]))
			row[j+1] = value
		}
		row[len(row)-1] = 9
		seafloor = append(seafloor, row)
	}
	seafloor = append(seafloor, zeroRowCopy)

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
