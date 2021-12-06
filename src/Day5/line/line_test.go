package line

import (
	"Day5/vertex"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		rowInfo string
	}
	tests := []struct {
		name string
		args args
		want Line
	}{
		// TODO: Add test cases.
		{
			name: "Standard pair",
			args: args{"0,9 -> 5,9"},
			want: Line{Start: vertex.Vertex{X: 0, Y: 9}, End: vertex.Vertex{X: 5, Y: 9}},
		},
		{
			name: "Standard pair",
			args: args{"8,0 -> 0,8"},
			want: Line{Start: vertex.Vertex{X: 8, Y: 0}, End: vertex.Vertex{X: 0, Y: 8}},
		},
		{
			name: "Standard pair",
			args: args{"9,4 -> 3,4"},
			want: Line{Start: vertex.Vertex{X: 9, Y: 4}, End: vertex.Vertex{X: 3, Y: 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.rowInfo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine_PointsBetweeen(t *testing.T) {
	type fields struct {
		Start vertex.Vertex
		End   vertex.Vertex
	}
	tests := []struct {
		name   string
		fields fields
		want   []vertex.Vertex
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Line{
				Start: tt.fields.Start,
				End:   tt.fields.End,
			}
			if got := l.PointsBetweeen(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Line.PointsBetweeen() = %v, want %v", got, tt.want)
			}
		})
	}
}
