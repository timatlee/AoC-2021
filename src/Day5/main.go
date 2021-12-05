package main

import (
	"Day5/line"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//filecontent := readfile("testinput.txt")
	filecontent := readfile("day5input.txt")

	// Get all the lines from the content, add to the lines slice
	var lines []line.Line
	for _, linecontent := range filecontent {
		l := line.New(linecontent)
		lines = append(lines, l)
	}

	// Find the max Y and Y so we can define our seafloor map.
	_, _, absmax := findMaxInLines(lines)

	// Now, we make a fixed size array of ints representing the floor.
	seafloor := make([][]int, absmax+1)
	for i := 0; i < len(seafloor); i++ {
		seafloor[i] = make([]int, absmax+1)
	}

	// Start drawing lines.
	for _, l := range lines {
		drawLine(seafloor, l)
	}

	// Count how many vertexes have more than two lines overlappin
	var overlapCount int = 0
	for x := 0; x < len(seafloor); x++ {
		for y := 0; y < len(seafloor[x]); y++ {
			if seafloor[x][y] >= 2 {
				overlapCount++
			}
		}
	}

	//fmt.Println(seafloor)
	fmt.Println(overlapCount)

}

func drawLine(s [][]int, l line.Line) {
	for _, v := range l.PointsBetweeen() {
		// This has to be "backwards" to our usual X, Y because we're defining all the rows FIRST (the Y coordinate), then which value of the COLUMN next (the X coordinate).
		s[v.Y][v.X] += 1
	}
}

func findMaxInLines(lines []line.Line) (int, int, int) {
	var xmax int
	var ymax int
	var absmax int

	for _, value := range lines {
		if value.Start.X > xmax {
			xmax = value.Start.X
		}
		if value.End.X > xmax {
			xmax = value.End.X
		}
		if value.Start.Y > ymax {
			ymax = value.Start.Y
		}
		if value.End.Y > ymax {
			ymax = value.End.Y
		}
	}
	if xmax > ymax {
		absmax = xmax
	} else {
		absmax = ymax
	}

	return xmax, ymax, absmax
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
