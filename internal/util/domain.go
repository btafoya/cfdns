package util

import (
	"fmt"

	"golang.org/x/net/publicsuffix"
)

func ExtractZone(hostname string) (string, error) {
	eTLD1, err := publicsuffix.EffectiveTLDPlusOne(hostname)
	if err != nil {
		return "", fmt.Errorf("zone extraction failed for %q: %w", hostname, err)
	}
	return eTLD1, nil
}
