package vertex

import (
	"reflect"
	"testing"
)

func TestNewFromInt(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want Vertex
	}{
		// TODO: Add test cases.
		{
			name: "Standard coordinates",
			args: args{1, 1},
			want: Vertex{X: 1, Y: 1},
		},
		{
			name: "Standard coordinates",
			args: args{0, 0},
			want: Vertex{X: 0, Y: 0},
		},
		{
			name: "Standard coordinates",
			args: args{100, 100},
			want: Vertex{X: 100, Y: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromInt(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromString(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name string
		args args
		want Vertex
	}{
		// TODO: Add test cases.
		{
			name: "Standard coordinates",
			args: args{"1", "1"},
			want: Vertex{X: 1, Y: 1},
		},
		{
			name: "Standard coordinates",
			args: args{"0", "0"},
			want: Vertex{X: 0, Y: 0},
		},
		{
			name: "Standard coordinates",
			args: args{"100", "100"},
			want: Vertex{X: 100, Y: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromString(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
