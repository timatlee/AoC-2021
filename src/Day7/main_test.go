package main

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_findMinValueInArray(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
		{
			name:  "Testing 1 value in array",
			args:  args{[]int{1}},
			want:  1,
			want1: 0,
		},
		{
			name:  "TEsting a few values in the array",
			args:  args{[]int{1, 2}},
			want:  1,
			want1: 0,
		},
		{
			name:  "Testing min value not in position 0",
			args:  args{[]int{2, 1}},
			want:  1,
			want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findMinValueInArray(tt.args.s)
			if got != tt.want {
				t.Errorf("findMinValueInArray() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findMinValueInArray() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_addBetweenRanges(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Add two adjacent numbers",
			args: args{a: 0, b: 1},
			want: 1,
		},
		{
			name: "Add non-adjacent numbers",
			args: args{a: 0, b: 3},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBetweenRanges(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBetweenRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readfile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readfile(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
