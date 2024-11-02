package aoc

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAocApi(t *testing.T) {
	require.Equal(t, AocApi().aocURL, aocUrl)
	require.Equal(t, AocApi().aocYear, aocYear)
}

func TestGetInput(t *testing.T) {
	cases := []struct {
		name       string
		day        int
		cookie     string
		statusCode int
		body       string
		wantBody   []byte
		wantErr    error
	}{
		{
			name:       "Bad status code",
			day:        1,
			cookie:     "test_cookie_value",
			statusCode: http.StatusInternalServerError,
			wantErr:    ErrFailedAPICall,
		},
		{
			name:       "Returns a Response",
			day:        1,
			cookie:     "test_cookie_value",
			statusCode: http.StatusOK,
			body:       `test`,
			wantBody:   []byte(`test`),
		},
		{
			name:    "Day is 0",
			day:     0,
			cookie:  "test_cookie_value",
			wantErr: ErrFailedAPICall,
		},
		{
			name:    "Day is out of range",
			day:     26,
			cookie:  "test_cookie_value",
			wantErr: ErrFailedAPICall,
		},
		{
			name:    "Empty cookie",
			day:     1,
			cookie:  "",
			wantErr: ErrFailedAPICall,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := fmt.Sprintf("/2022/day/%d/input", tc.day)
				if r.URL.Path != expectedPath {
					t.Errorf("expected URL path %s, got %s", expectedPath, r.URL.Path)
				}

				cookie, err := r.Cookie("session")
				if err != nil || cookie.Value != tc.cookie {
					t.Errorf("expected session cookie with value 'test_cookie_value', got %v", cookie)
				}

				w.WriteHeader(tc.statusCode)
				w.Write([]byte(tc.body))
			}))
			defer testServer.Close()

			aocClient := NewAoc(testServer.URL, 2022)

			gotResponse, gotErr := aocClient.GetInput(tc.day, tc.cookie)

			require.ErrorIs(t, gotErr, tc.wantErr)
			require.Equal(t, gotResponse, tc.wantBody)
		})
	}
}
