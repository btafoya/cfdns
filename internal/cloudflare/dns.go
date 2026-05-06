package cloudflare

import (
	"context"
	"fmt"

	cf "github.com/cloudflare/cloudflare-go"
)

func FindRecord(ctx context.Context, api *cf.API, zoneID, hostname, recType string) (*cf.DNSRecord, error) {
	records, _, err := api.ListDNSRecords(ctx, cf.ZoneIdentifier(zoneID), cf.ListDNSRecordsParams{
		Name: hostname,
		Type: recType,
	})
	if err != nil {
		return nil, fmt.Errorf("DNS lookup failed: %w", err)
	}
	if len(records) > 0 {
		return &records[0], nil
	}
	return nil, nil
}

func CreateRecord(ctx context.Context, api *cf.API, zoneID, recType, hostname, content string, ttl int, proxied bool) error {
	_, err := api.CreateDNSRecord(ctx, cf.ZoneIdentifier(zoneID), cf.CreateDNSRecordParams{
		Type:    recType,
		Name:    hostname,
		Content: content,
		TTL:     ttl,
		Proxied: cf.BoolPtr(proxied),
	})
	if err != nil {
		return fmt.Errorf("create failed: %w", err)
	}
	return nil
}

func UpdateRecord(ctx context.Context, api *cf.API, zoneID, recordID, recType, hostname, content string, ttl int, proxied bool) error {
	_, err := api.UpdateDNSRecord(ctx, cf.ZoneIdentifier(zoneID), cf.UpdateDNSRecordParams{
		ID:      recordID,
		Type:    recType,
		Name:    hostname,
		Content: content,
		TTL:     ttl,
		Proxied: cf.BoolPtr(proxied),
	})
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}
	return nil
}

func DeleteRecord(ctx context.Context, api *cf.API, zoneID, recordID string) error {
	if err := api.DeleteDNSRecord(ctx, cf.ZoneIdentifier(zoneID), recordID); err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}
	return nil
}

func ListRecords(ctx context.Context, api *cf.API, zoneID string) ([]cf.DNSRecord, error) {
	records, _, err := api.ListDNSRecords(ctx, cf.ZoneIdentifier(zoneID), cf.ListDNSRecordsParams{})
	if err != nil {
		return nil, fmt.Errorf("list failed: %w", err)
	}
	return records, nil
}
