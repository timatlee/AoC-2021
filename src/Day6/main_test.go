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

func Test_growFish(t *testing.T) {
	type args struct {
		fish []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "One fish, initial timer of 3, day: 1",
			args: args{[]int{0, 0, 1, 0, 0, 0, 0, 0, 0}},
			want: []int{0, 1, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "One fish, initial timer of 3, day: 2",
			args: args{[]int{0, 1, 0, 0, 0, 0, 0, 0, 0}},
			want: []int{1, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "One fish, initial timer of 3, day: 3",
			args: args{[]int{1, 0, 0, 0, 0, 0, 0, 0, 0}},
			want: []int{0, 0, 0, 0, 0, 0, 1, 0, 1},
		},
		{
			name: "One fish, initial timer of 3, day: 4",
			args: args{[]int{0, 0, 0, 0, 0, 0, 1, 0, 1}},
			want: []int{0, 0, 0, 0, 0, 1, 0, 1, 0},
		},
		{
			name: "One fish, initial timer of 3, day: 5",
			args: args{[]int{0, 0, 0, 0, 0, 1, 0, 1, 0}},
			want: []int{0, 0, 0, 0, 1, 0, 1, 0, 0},
		},
		{
			name: "One fish, initial timer of 3, day: 6",
			args: args{[]int{0, 0, 0, 0, 1, 0, 1, 0, 0}},
			want: []int{0, 0, 0, 1, 0, 1, 0, 0, 0},
		},
		{
			name: "One fish, initial timer of 3, day: 7",
			args: args{[]int{0, 0, 0, 1, 0, 1, 0, 0, 0}},
			want: []int{0, 0, 1, 0, 1, 0, 0, 0, 0},
		},
		{
			name: "One fish, initial timer of 3, day: 8",
			args: args{[]int{0, 0, 1, 0, 1, 0, 0, 0, 0}},
			want: []int{0, 1, 0, 1, 0, 0, 0, 0, 0},
		},
		{
			name: "One fish, initial timer of 3, day: 9",
			args: args{[]int{0, 1, 0, 1, 0, 0, 0, 0, 0}},
			want: []int{1, 0, 1, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "One fish, initial timer of 3, day: 10",
			args: args{[]int{1, 0, 1, 0, 0, 0, 0, 0, 0}},
			want: []int{0, 1, 0, 0, 0, 0, 1, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := growFish(tt.args.fish); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("growFish() = %v, want %v", got, tt.want)
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

func Test_sumArray(t *testing.T) {
	type args struct {
		fish []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Sum with 1 entry",
			args: args{[]int{1}},
			want: 1,
		},
		{
			name: "Sum with multiple entries",
			args: args{[]int{1, 2}},
			want: 3,
		},
		{
			name: "Sum with a lot of entries",
			args: args{[]int{3, 5, 3, 2, 2, 1, 5, 1, 4}},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumArray(tt.args.fish); got != tt.want {
				t.Errorf("sumArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
