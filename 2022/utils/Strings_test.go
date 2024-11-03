package utils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommonChars(t *testing.T) {
	tests := []struct {
		name       string
		str1       string
		str2       string
		wantCommon []rune
	}{
		{
			name:       "No common characters",
			str1:       "abc",
			str2:       "def",
			wantCommon: []rune{},
		},
		{
			name:       "Differentiating case",
			str1:       "abc",
			str2:       "Abc",
			wantCommon: []rune{'b'},
		},
		{
			name:       "Multiple common characters",
			str1:       "hello",
			str2:       "world",
			wantCommon: []rune{'l', 'o'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CommonChars(tt.str1, tt.str2)
			if reflect.DeepEqual(got, tt.wantCommon) {
				t.Errorf("CommonChars(%s, %s) = %v, want %v", tt.str1, tt.str2, got, tt.wantCommon)
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
			gotNumbers, _ := ExtractNumbers(tc.input)

			require.Len(t, gotNumbers, len(tc.wantNumbers))
			if reflect.DeepEqual(gotNumbers, tc.wantNumbers) {
				t.Errorf("extractNumbers(%s) = %v, want %v", tc.input, gotNumbers, tc.wantNumbers)
			}
		})
	}
}

func TestDuplicateChars(t *testing.T) {
	cases := []struct {
		name      string
		input     string
		wantChars map[rune]int
	}{
		{
			name:      "No duplicates found",
			input:     "abc",
			wantChars: map[rune]int{'c': 1, 'b': 1, 'a': 1},
		},
		{
			name:      "Duplicates found",
			input:     "abcadc",
			wantChars: map[rune]int{'a': 2, 'b': 1, 'c': 2, 'd': 1},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotChars := DuplicateChars(tc.input)

			require.Len(t, gotChars, len(tc.wantChars))

			for k, v := range tc.wantChars {
				if gotChars[k] != v {
					t.Errorf("DuplicateChars(%s) = %v, want %v", tc.input, gotChars, tc.wantChars)
				}
			}
		})
	}
}

func TestHasDuplicateChars(t *testing.T) {
	cases := []struct {
		name           string
		input          string
		wantDuplicates bool
	}{
		{
			name:           "No duplicates found",
			input:          "abc",
			wantDuplicates: false,
		},
		{
			name:           "Duplicates found",
			input:          "abcadc",
			wantDuplicates: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.wantDuplicates, HasDuplicateChars(tc.input))
		})
	}
}
