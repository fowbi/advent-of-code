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
