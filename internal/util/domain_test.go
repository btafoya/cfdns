package util

import "testing"

func TestExtractZone(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"example.com", "example.com"},
		{"sub.example.com", "example.com"},
		{"sub.sub.example.com", "example.com"},
		{"example.co.uk", "example.co.uk"},
		{"test.example.co.uk", "example.co.uk"},
	}
	for _, tt := range tests {
		got, err := ExtractZone(tt.input)
		if err != nil {
			t.Errorf("ExtractZone(%q) unexpected error: %v", tt.input, err)
			continue
		}
		if got != tt.expected {
			t.Errorf("ExtractZone(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}
