package main

import (
	_ "embed"
	"testing"
)

func TestPart01(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "Sample input #1",
			input: `...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`,
			want: 2,
		},
		{
			name: "Sample input #2",
			input: `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`,
			want: 4,
		},
		{
			name: "Sample input #3",
			input: `10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01`,
			want: 3,
		},
		{
			name: "Sample input #4",
			input: `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
			want: 36,
		},
		{
			name:  "Actual input",
			input: input,
			want:  1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := part_01(tc.input); got != tc.want {
				t.Errorf("part_01 = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestPart02(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "Sample input",
			input: `.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....`,
			want: 3,
		},
		{
			name: "Sample input #2",
			input: `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`,
			want: 13,
		},
		{
			name: "Sample input #3",
			input: `012345
123456
234567
345678
4.6789
56789`,
			want: 227,
		},
		{
			name: "Sample input #4",
			input: `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
			want: 81,
		},
		{
			name:  "Actual input",
			input: input,
			want:  1609,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := part_02(tc.input); got != tc.want {
				t.Errorf("part_02 = %v, want %v", got, tc.want)
			}
		})
	}
}
