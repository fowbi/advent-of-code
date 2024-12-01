package utils

import "testing"

func TestFindLargestNumber(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{
			name:     "Mixed positive numbers",
			numbers:  []int{45, 22, 68, 90, 12},
			expected: 90,
		},
		{
			name:     "All negative numbers",
			numbers:  []int{-5, -23, -1, -55},
			expected: -1,
		},
		{
			name:     "All zeros",
			numbers:  []int{0, 0, 0, 0, 0},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindLargestNumber(tt.numbers)
			if got != tt.expected {
				t.Errorf("findLargestNumber(%v) = %v, want %v", tt.numbers, got, tt.expected)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	tests := []struct {
		name     string
		number1  int
		number2  int
		expected int
	}{
		{
			name:     "Zero",
			number1:  1,
			number2:  1,
			expected: 0,
		},
		{
			name:     "First number larger than second",
			number1:  3,
			number2:  1,
			expected: 2,
		},
		{
			name:     "Second number larger than first",
			number1:  1,
			number2:  3,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Diff(tt.number1, tt.number2)
			if got != tt.expected {
				t.Errorf("Diff(%v, %v) = %v, want %v", tt.number1, tt.number2, got, tt.expected)
			}
		})
	}
}
