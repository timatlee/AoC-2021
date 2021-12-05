package line

import (
	"Day5/vertex"
	"regexp"
)

type Line struct {
	Start vertex.Vertex
	End   vertex.Vertex
}

func New(rowInfo string) Line {
	var compiledRegex = regexp.MustCompile(`^(?P<x1>\d+),(?P<y1>\d+) -> (?P<x2>\d+),(?P<y2>\d+)$`)
	matches := compiledRegex.FindStringSubmatch(rowInfo)
	x1 := matches[compiledRegex.SubexpIndex("x1")]
	y1 := matches[compiledRegex.SubexpIndex("y1")]
	x2 := matches[compiledRegex.SubexpIndex("x2")]
	y2 := matches[compiledRegex.SubexpIndex("y2")]

	l := Line{
		Start: vertex.NewFromString(x1, y1),
		End:   vertex.NewFromString(x2, y2),
	}

	return l
}

func (l *Line) IsStraight() bool {
	return (l.Start.X == l.End.X) || (l.Start.Y == l.End.Y)
}
