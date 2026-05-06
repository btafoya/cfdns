package cloudflare

import (
	"fmt"
	"os"

	cf "github.com/cloudflare/cloudflare-go"
)

func NewClient() (*cf.API, error) {
	token := os.Getenv("CLOUDFLARE_API_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("CLOUDFLARE_API_TOKEN not set")
	}
	return cf.NewWithAPIToken(token)
}
