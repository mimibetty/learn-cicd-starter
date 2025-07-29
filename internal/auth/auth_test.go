package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name:          "No authorization header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Empty authorization header",
			headers: http.Header{
				"Authorization": []string{""},
			},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed authorization header - missing API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name: "Malformed authorization header - wrong prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer some-token"},
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name: "Valid authorization header",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-secret-api-key"},
			},
			expectedKey:   "my-secret-api-key",
			expectedError: nil,
		},
		{
			name: "Valid authorization header with extra spaces",
			headers: http.Header{
				"Authorization": []string{"ApiKey key-with-extra-parts more-data"},
			},
			expectedKey:   "key-with-extra-parts",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			if key != tt.expectedKey {
				t.Errorf("GetAPIKey() key = %v, want %v", key, tt.expectedKey)
			}

			if tt.expectedError == nil {
				if err != nil {
					t.Errorf("GetAPIKey() error = %v, want nil", err)
				}
			} else {
				if err == nil {
					t.Errorf("GetAPIKey() error = nil, want %v", tt.expectedError)
				} else if err.Error() != tt.expectedError.Error() {
					t.Errorf("GetAPIKey() error = %v, want %v", err, tt.expectedError)
				}
			}
		})
	}
}

func TestGetAPIKey_CaseInsensitiveHeader(t *testing.T) {
	// Test that header names are case-insensitive (standard HTTP behavior)
	headers := http.Header{
		"authorization": []string{"ApiKey test-key-lowercase"},
	}

	key, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("GetAPIKey() with lowercase header failed: %v", err)
	}
	if key != "test-key-lowercase" {
		t.Errorf("GetAPIKey() key = %v, want %v", key, "test-key-lowercase")
	}
}
