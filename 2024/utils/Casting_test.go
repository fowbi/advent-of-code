package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCastStringToInt(t *testing.T) {
	faultyCases := []struct {
		name        string
		faultyValue string
	}{
		{
			name:        "It can not convert a boolean to an integer",
			faultyValue: "true",
		},
		{
			name:        "It can not convert a array to an integer",
			faultyValue: "[]",
		},
	}
	cases := []struct {
		name      string
		value     string
		wantValue int
	}{
		{
			name:      "It can convert a string to an integer",
			value:     "123",
			wantValue: 123,
		},
	}

	for _, tc := range faultyCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code did not panic")
				}
			}()

			CastStringToInt(tc.faultyValue)
		})
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.wantValue, CastStringToInt(tc.value))
		})
	}
}
