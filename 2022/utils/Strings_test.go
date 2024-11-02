package utils

import (
	"reflect"
	"testing"
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
