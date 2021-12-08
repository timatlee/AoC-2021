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

func Test_splitSignalsAndDigits(t *testing.T) {
	type args struct {
		l string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		// TODO: Add test cases.
		{
			name:  "Testing sample input 1",
			args:  args{l: "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe"},
			want:  []string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"},
			want1: []string{"abcdefg", "bcdef", "bcdefg", "bceg"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitSignalsAndDigits(tt.args.l)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitSignalsAndDigits() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("splitSignalsAndDigits() got1 = %v, want %v", got1, tt.want1)
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

func Test_mapSegcountToDigit(t *testing.T) {
	type args struct {
		s int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "2 segements must be a 1.",
			args: args{s: 2},
			want: 1,
		},
		{
			name: "4 segements must be a 4.",
			args: args{s: 4},
			want: 4,
		},
		{
			name: "3 segments must be a 7",
			args: args{s: 3},
			want: 7,
		},
		{
			name: "7 segments must be an 8",
			args: args{s: 7},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapSegcountToDigit(tt.args.s); got != tt.want {
				t.Errorf("mapSegcountToDigit() = %v, want %v", got, tt.want)
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

func Test_determine_letter_to_number(t *testing.T) {
	type args struct {
		input []string
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
			if got := determine_letter_to_number(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("determine_letter_to_number() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_difference(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Simple test",
			args: args{"abc", "bc"},
			want: "a",
		},
		{
			name: "Empty result",
			args: args{"abc", "abc"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := difference(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_find_match_in_input(t *testing.T) {
	type args struct {
		one   string
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Simple Working true test",
			args: args{one: "ab", input: "abc"},
			want: true,
		},
		{
			name: "Simple working false test",
			args: args{one: "ab", input: "bcd"},
			want: false,
		},
		{
			name: "Simple working false test",
			args: args{one: "ab", input: "cde"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find_match_in_input(tt.args.one, tt.args.input); got != tt.want {
				t.Errorf("find_match_in_input() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexOfArray(t *testing.T) {
	type args struct {
		s    []string
		item string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Simple Test",
			args: args{s: []string{"a", "b", "c"}, item: "a"},
			want: 0,
		},
		{
			name: "Simple Test non-zero index",
			args: args{s: []string{"a", "b", "c"}, item: "b"},
			want: 1,
		},
		{
			name: "Simple Test not found item",
			args: args{s: []string{"a", "b", "c"}, item: "d"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := indexOfArray(tt.args.s, tt.args.item); got != tt.want {
				t.Errorf("indexOfArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
