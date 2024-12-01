package main

import (
	_ "embed"
	"testing"
)

//go:embed sample-input.txt
var sampleInput string

func TestPart01(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Sample input",
			input: sampleInput,
			want:  11,
		},
		{
			name:  "Actual input",
			input: input,
			want:  1580061,
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
			name:  "Sample input",
			input: sampleInput,
			want:  31,
		},
		{
			name:  "Actual input",
			input: input,
			want:  23046913,
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
