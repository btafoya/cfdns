package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/btafoya/cfdns/internal/cli"
	cfapi "github.com/btafoya/cfdns/internal/cloudflare"
	"github.com/btafoya/cfdns/internal/util"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list <domain>",
	Short: "List DNS records for a domain",
	Args:  cobra.ExactArgs(1),
	RunE:  runList,
}

func runList(cmd *cobra.Command, args []string) error {
	domain := args[0]

	if err := util.ValidateHostname(domain); err != nil {
		cli.PrintError(err.Error(), flagJSON)
		os.Exit(2)
	}

	api, err := cfapi.NewClient()
	if err != nil {
		cli.PrintError(err.Error(), flagJSON)
		os.Exit(1)
	}

	ctx := context.Background()

	zone, err := util.ExtractZone(domain)
	if err != nil {
		cli.PrintError(err.Error(), flagJSON)
		os.Exit(2)
	}

	zoneID, err := cfapi.GetZoneID(ctx, api, zone)
	if err != nil {
		cli.PrintError(err.Error(), flagJSON)
		os.Exit(3)
	}

	records, err := cfapi.ListRecords(ctx, api, zoneID)
	if err != nil {
		cli.PrintError(err.Error(), flagJSON)
		os.Exit(3)
	}

	if flagJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		_ = enc.Encode(records)
		return nil
	}

	if len(records) == 0 {
		fmt.Println("No records found.")
		return nil
	}

	for _, r := range records {
		proxied := ""
		if r.Proxied != nil && *r.Proxied {
			proxied = " [proxied]"
		}
		fmt.Printf("%-6s %-40s %s%s\n", r.Type, r.Name, r.Content, proxied)
	}
	return nil
}
