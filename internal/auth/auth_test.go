package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		headers        http.Header
		expectedAPIKey string
		expectedError  error
	}{
		{
			headers: http.Header{
				"Authorization": []string{"ApiKey 1234"},
			},
			expectedAPIKey: "1234",
			expectedError:  nil,
		},
		{
			headers:        http.Header{},
			expectedAPIKey: "",
			expectedError:  ErrNoAuthHeaderIncluded,
		},
		{
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedAPIKey: "",
			expectedError:  ErrMalformedAuthHeader,
		},
	}

	for _, c := range cases {
		apiKey, err := GetAPIKey(c.headers)

		if !errors.Is(err, c.expectedError) {
			t.Errorf("Expected no error, got %v", err)
		}

		if apiKey != c.expectedAPIKey {
			t.Errorf("Expected %s, got %s", c.expectedAPIKey, apiKey)
		}
	}
}

