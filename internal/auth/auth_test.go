package auth

import (
	"net/http"
	"testing"
)

// Test when there is no Authorization header
func TestGetAPIKey_NoHeader(t *testing.T) {
	headers := http.Header{}

	key, err := GetAPIKey(headers)

	if key != "" {
		t.Errorf("Expected empty key, got %s", key)
	}
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}

// Test when Authorization header has correct format
func TestGetAPIKey_ValidHeader(t *testing.T) {
	headers := http.Header{
		"Authorization": []string{"ApiKey my-secret-key"},
	}

	key, err := GetAPIKey(headers)

	if key != "my-secret-key" {
		t.Errorf("Expected 'my-secret-key', got %s", key)
	}
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
