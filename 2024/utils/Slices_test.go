package utils

import (
	"reflect"
	"testing"
)

func TestRemove(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		index    int
		expected []int
	}{
		{
			name:     "Remove first element",
			slice:    []int{1, 2, 3, 4},
			index:    0,
			expected: []int{2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Remove(tt.slice, tt.index)
			if reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Remove(%v, %v) = %v, want %v", tt.slice, tt.index, got, tt.expected)
			}
		})
	}
}
