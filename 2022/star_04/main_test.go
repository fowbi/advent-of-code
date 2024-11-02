package main

import (
	_ "embed"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
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
			want:  2,
		},
		{
			name:  "Actual input",
			input: input,
			want:  569,
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
			want:  4,
		},
		{
			name:  "Actual input",
			input: input,
			want:  936,
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

func TestExtractNumbers(t *testing.T) {
	cases := []struct {
		name        string
		input       string
		wantNumbers []int
	}{
		{
			name:        "Sample input",
			input:       "1-2,3-4",
			wantNumbers: []int{4, 3, 2, 1},
		},
		{
			name:        "Actual input",
			input:       "abc123",
			wantNumbers: []int{321},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotNumbers, _ := extractNumbers(tc.input)

			require.Len(t, gotNumbers, len(tc.wantNumbers))
			if reflect.DeepEqual(gotNumbers, tc.wantNumbers) {
				t.Errorf("extractNumbers(%s) = %v, want %v", tc.input, gotNumbers, tc.wantNumbers)
			}
		})
	}
}

func TestFullyOverlap(t *testing.T) {
	cases := []struct {
		name        string
		first       Pair
		second      Pair
		wantOverlap bool
	}{
		{
			name:        "First overlaps second",
			first:       Pair{x: 1, y: 10},
			second:      Pair{x: 3, y: 4},
			wantOverlap: true,
		},
		{
			name:        "Second overlaps first",
			first:       Pair{x: 3, y: 4},
			second:      Pair{x: 1, y: 10},
			wantOverlap: true,
		},
		{
			name:        "Overlaps but not fully",
			first:       Pair{x: 3, y: 5},
			second:      Pair{x: 5, y: 10},
			wantOverlap: false,
		},
		{
			name:        "Does not overlap",
			first:       Pair{x: 3, y: 4},
			second:      Pair{x: 8, y: 10},
			wantOverlap: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			sap := SectionAssignmentPairs{first: tc.first, second: tc.second}

			require.Equal(t, tc.wantOverlap, sap.FullyOverlap())
		})
	}
}

func TestOverlap(t *testing.T) {
	cases := []struct {
		name        string
		first       Pair
		second      Pair
		wantOverlap bool
	}{
		{
			name:        "First overlaps second",
			first:       Pair{x: 1, y: 10},
			second:      Pair{x: 3, y: 4},
			wantOverlap: true,
		},
		{
			name:        "Second overlaps first",
			first:       Pair{x: 3, y: 4},
			second:      Pair{x: 1, y: 10},
			wantOverlap: true,
		},
		{
			name:        "Overlaps but not fully",
			first:       Pair{x: 3, y: 5},
			second:      Pair{x: 5, y: 10},
			wantOverlap: true,
		},
		{
			name:        "Does not overlap",
			first:       Pair{x: 3, y: 4},
			second:      Pair{x: 8, y: 10},
			wantOverlap: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			sap := SectionAssignmentPairs{first: tc.first, second: tc.second}

			require.Equal(t, tc.wantOverlap, sap.Overlap())
		})
	}
}
