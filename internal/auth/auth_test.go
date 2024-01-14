package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_Success(t *testing.T) {
	headers := http.Header{}
	headers.Set("authorization", "ApiKey 12345")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if apiKey != "12345" {
		t.Errorf("Expected apiKey to be '12345', got '%s'", apiKey)
	}
}

func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected error '%v', got '%v'", ErrNoAuthHeaderIncluded, err)
	}
}

func TestGetAPIKey_MalformedAuthHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "MalformedHeader")

	_, err := GetAPIKey(headers)
	if err == nil || err.Error() != "malformed authorization header" {
		t.Errorf("Expected error 'malformed authorization header', got '%v'", err)
	}
}
