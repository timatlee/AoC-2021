package vertex

import "strconv"

type Vertex struct {
	X, Y int
}

func NewFromInt(x, y int) Vertex {
	v := Vertex{X: x, Y: y}
	return v
}

func NewFromString(x, y string) Vertex {
	x_num, _ := strconv.Atoi(x)
	y_num, _ := strconv.Atoi(y)
	v := Vertex{X: x_num, Y: y_num}
	return v
}
