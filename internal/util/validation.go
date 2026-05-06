package util

import (
	"fmt"
	"net"
	"strings"
)

func ValidateHostname(hostname string) error {
	if hostname == "" {
		return fmt.Errorf("hostname required")
	}
	if !strings.Contains(hostname, ".") {
		return fmt.Errorf("invalid hostname: %q", hostname)
	}
	return nil
}

func ValidateIP(ip string) error {
	if net.ParseIP(ip) == nil {
		return fmt.Errorf("invalid IP address: %q", ip)
	}
	return nil
}

func IsIPv6(ip string) bool {
	parsed := net.ParseIP(ip)
	if parsed == nil {
		return false
	}
	return parsed.To4() == nil
}
