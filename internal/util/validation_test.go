package util

import "testing"

func TestValidateHostname(t *testing.T) {
	tests := []struct {
		input string
		valid bool
	}{
		{"example.com", true},
		{"sub.example.com", true},
		{"test.example.co.uk", true},
		{"", false},
		{"noDot", false},
	}
	for _, tt := range tests {
		err := ValidateHostname(tt.input)
		if tt.valid && err != nil {
			t.Errorf("expected valid for %q, got: %v", tt.input, err)
		}
		if !tt.valid && err == nil {
			t.Errorf("expected invalid for %q", tt.input)
		}
	}
}

func TestValidateIP(t *testing.T) {
	tests := []struct {
		input string
		valid bool
	}{
		{"192.168.1.1", true},
		{"2001:db8::1", true},
		{"not-an-ip", false},
		{"999.999.999.999", false},
	}
	for _, tt := range tests {
		err := ValidateIP(tt.input)
		if tt.valid && err != nil {
			t.Errorf("expected valid for %q, got: %v", tt.input, err)
		}
		if !tt.valid && err == nil {
			t.Errorf("expected invalid for %q", tt.input)
		}
	}
}

func TestIsIPv6(t *testing.T) {
	if IsIPv6("192.168.1.1") {
		t.Error("IPv4 should not be IPv6")
	}
	if !IsIPv6("2001:db8::1") {
		t.Error("IPv6 address should be detected")
	}
}
