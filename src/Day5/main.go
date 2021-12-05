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
	xmax, ymax := findMaxInLines(lines)

	// Now, we make a fixed size array of ints representing the floor.
	seafloor := make([][]int, xmax+10)
	for i := 0; i < len(seafloor); i++ {
		seafloor[i] = make([]int, ymax+10)
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

	fmt.Println(seafloor)
	fmt.Println(overlapCount)

}

func drawLine(s [][]int, l line.Line) {
	if l.IsStraight() {
		// But is it going horizontally or vertically?
		if l.Start.X == l.End.X {
			if l.Start.Y < l.End.Y {
				for i := l.Start.Y; i <= l.End.Y; i++ {
					s[i][l.Start.X] += 1
				}
			} else {
				for i := l.End.Y; i <= l.Start.Y; i++ {
					s[i][l.Start.X] += 1
				}
			}
		} else {
			if l.Start.X < l.End.X {
				for i := l.Start.X; i <= l.End.X; i++ {
					s[l.Start.Y][i] += 1
				}
			} else {
				for i := l.End.X; i <= l.Start.X; i++ {
					s[l.Start.Y][i] += 1
				}
			}
		}
	}
}

func findMaxInLines(lines []line.Line) (int, int) {
	var xmax int
	var ymax int

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

	return xmax, ymax
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
