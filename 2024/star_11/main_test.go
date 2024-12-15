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
			name:  "Sample input",
			input: "125 17",
			want:  55312,
		},
		{
			name:  "Actual input",
			input: input,
			want:  228668,
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
			input: "",
			want:  22938365706844,
		},
		{
			name:  "Actual input",
			input: input,
			want:  270673834779359,
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
