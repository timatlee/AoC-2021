package main

import (
	"fmt"

	aocutils "github.com/timatlee/AoC-Common"
)

type Vertex struct {
	X int
	Y int
}

type Instructions struct {
	Direction string
	Size      int
}

func main() {
	content := aocutils.Readfile("input.txt")
	vertexes, instructions := parseInput(content)

	// Part 1
	for i := range instructions[:1] {
		switch instructions[i].Direction {
		case "x":
			foldAlongX(vertexes, instructions[i].Size)
		case "y":
			foldAlongY(vertexes, instructions[i].Size)
		}
		//fmt.Printf("vertexes: %v\n", vertexes)
		fmt.Printf("len(vertexes): %v\n", len(vertexes))
	}

	vertexes, instructions = parseInput(content)

	// Part 2
	for i := range instructions {
		switch instructions[i].Direction {
		case "x":
			foldAlongX(vertexes, instructions[i].Size)
		case "y":
			foldAlongY(vertexes, instructions[i].Size)
		}
		//fmt.Printf("vertexes: %v\n", vertexes)
		fmt.Printf("len(vertexes): %v\n", len(vertexes))
	}
	printVertexes(vertexes)

}

func printVertexes(m map[Vertex]bool) {
	// Figure out X max
	// Figure out Y max
	var xs []int
	var ys []int
	for c := range m {
		xs = append(xs, c.X)
		ys = append(ys, c.Y)
	}
	_, xmax := aocutils.MinMaxInt(xs)
	_, ymax := aocutils.MinMaxInt(ys)

	for y := 0; y < ymax+1; y++ {
		for x := 0; x < xmax+1; x++ {
			if m[Vertex{X: x, Y: y}] == true {
				fmt.Print(".")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	// Repeat 0 to Ymax
	// Start drawing points between 0 and Xmax
}

func foldAlongX(m map[Vertex]bool, c int) {
	for v := range m {
		if v.X > c {
			oldY := v.Y
			newX := c - (v.X - c)
			delete(m, v)
			m[Vertex{X: newX, Y: oldY}] = true
		}
	}
}

func foldAlongY(m map[Vertex]bool, c int) {
	for v := range m {
		if v.Y > c {
			oldX := v.X
			newY := c - (v.Y - c)
			delete(m, v)
			m[Vertex{X: oldX, Y: newY}] = true
		}
	}
}

func parseInput(s []string) (map[Vertex]bool, []Instructions) {
	coords := map[Vertex]bool{}
	instructions := []Instructions{}

	for i := range s {
		if len(s[i]) > 0 {
			var x, y int
			_, err := fmt.Sscanf(s[i], "%d,%d", &x, &y)
			if nil == err {
				coords[Vertex{X: x, Y: y}] = true
			}

			var size int
			var dir string
			_, err = fmt.Sscanf(s[i], "fold along %1s=%d", &dir, &size)
			if nil == err {
				instructions = append(instructions, Instructions{Direction: string(dir), Size: size})
			}
		}
	}

	return coords, instructions

}
