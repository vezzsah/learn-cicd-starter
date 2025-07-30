package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		wantedToken string
		wantedError bool
	}{
		{
			name: "Valid ApiKey token",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid_token"},
			},
			wantedToken: "valid_token",
			wantedError: false,
		},
		{
			name: "Invalid ApiKey token prefix",
			headers: http.Header{
				"Authorization": []string{"InvalidKey valid_token"},
			},
			wantedToken: "",
			wantedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantedError {
				t.Errorf("GetAPIKey() error = %v, wantedError = %v", err, tt.wantedError)
				return
			}
			if gotToken != tt.wantedToken {
				t.Errorf("GetAPIKey() gotToken = %v, want = %v", gotToken, tt.wantedToken)
			}
		})
	}
}
