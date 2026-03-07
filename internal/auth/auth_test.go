package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		authHeader  string
		expectedKey string
		expectError bool
	}{
		{
			name:        "missing header",
			authHeader:  "",
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "malformed header no key",
			authHeader:  "ApiKey",
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "wrong scheme",
			authHeader:  "Bearer abc123",
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "valid api key",
			authHeader:  "ApiKey abc123",
			expectedKey: "abc123",
			expectError: false,
		},
		{
			name:        "valid api key with spaces",
			authHeader:  "ApiKey my-secret-key",
			expectedKey: "my-secret-key",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.authHeader != "" {
				headers.Set("Authorization", tt.authHeader)
			}

			key, err := GetAPIKey(headers)

			if tt.expectError && err == nil {
				t.Fatalf("expected error but got none")
			}

			if !tt.expectError && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if key != tt.expectedKey {
				t.Fatalf("expected key %q, got %q", tt.expectedKey, key)
			}
		})
	}
}
