package cloudflare

import (
	"context"
	"fmt"

	cf "github.com/cloudflare/cloudflare-go"
)

func GetZoneID(ctx context.Context, api *cf.API, zoneName string) (string, error) {
	resp, err := api.ListZonesContext(ctx, cf.WithZoneFilters(zoneName, "", ""))
	if err != nil {
		return "", fmt.Errorf("zone lookup failed: %w", err)
	}
	if len(resp.Result) == 0 {
		return "", fmt.Errorf("zone not found: %s", zoneName)
	}
	return resp.Result[0].ID, nil
}
