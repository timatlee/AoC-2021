package line

import (
	"Day5/vertex"
	"fmt"
)

type Line struct {
	Start vertex.Vertex
	End   vertex.Vertex
}

func New(rowInfo string) Line {
	var x1, y1, x2, y2 int
	fmt.Sscanf(rowInfo, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

	l := Line{
		Start: vertex.NewFromInt(x1, y1),
		End:   vertex.NewFromInt(x2, y2),
	}

	return l
}

func (l *Line) PointsBetweeen() []vertex.Vertex {
	var vs []vertex.Vertex

	// Horizontal starigth line
	if l.Start.X == l.End.X {
		if l.Start.Y < l.End.Y {
			for y := l.Start.Y; y <= l.End.Y; y++ {
				vs = append(vs, vertex.NewFromInt(l.Start.X, y))
			}
		} else {
			for y := l.End.Y; y <= l.Start.Y; y++ {
				vs = append(vs, vertex.NewFromInt(l.Start.X, y))
			}
		}
	} else // Vertical straigth line.
	if l.Start.Y == l.End.Y {
		if l.Start.X < l.End.X {
			for x := l.Start.X; x <= l.End.X; x++ {
				vs = append(vs, vertex.NewFromInt(x, l.Start.Y))
			}
		} else {
			for x := l.End.X; x <= l.Start.X; x++ {
				vs = append(vs, vertex.NewFromInt(x, l.Start.Y))
			}
		}
	} else // Angled line.
	{
		// Find slope, but since origin is top-left, not bottom-left, invert this?
		m := -1 * ((l.End.Y - l.Start.Y) / (l.End.X - l.Start.X))
		if m < 0 {
			if l.Start.X < l.End.X {
				for x := l.Start.X; x <= l.End.X; x++ {
					vs = append(vs, vertex.NewFromInt(x, l.Start.Y+(x-l.Start.X)))
				}
			} else {
				for x := l.End.X; x <= l.Start.X; x++ {
					vs = append(vs, vertex.NewFromInt(x, l.End.Y+(x-l.End.X)))
				}
			}
		} else {
			if l.Start.X < l.End.X {
				for x := l.Start.X; x <= l.End.X; x++ {
					vs = append(vs, vertex.NewFromInt(x, l.Start.Y-(x-l.Start.X)))
				}
			} else {
				for x := l.End.X; x <= l.Start.X; x++ {
					vs = append(vs, vertex.NewFromInt(x, l.End.Y-(x-l.End.X)))
				}
			}
		}
	}

	return vs
}
