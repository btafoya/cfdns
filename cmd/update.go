package cmd

import (
	"context"
	"os"

	"github.com/btafoya/cfdns/internal/cli"
	cfapi "github.com/btafoya/cfdns/internal/cloudflare"
	"github.com/btafoya/cfdns/internal/util"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update <hostname> <content>",
	Short: "Update an existing DNS record",
	Args:  cobra.ExactArgs(2),
	RunE:  runUpdate,
}

func runUpdate(cmd *cobra.Command, args []string) error {
	hostname, content := args[0], args[1]

	if err := util.ValidateHostname(hostname); err != nil {
		cli.PrintError(err.Error(), flagJSON)
		os.Exit(2)
	}

	recType := resolveType(content)

	if recType == "A" || recType == "AAAA" {
		if err := util.ValidateIP(content); err != nil {
			cli.PrintError(err.Error(), flagJSON)
			os.Exit(2)
		}
	}

	api, err := cfapi.NewClient()
	if err != nil {
		cli.PrintError(err.Error(), flagJSON)
		os.Exit(1)
	}

	ctx := context.Background()

	zone, err := util.ExtractZone(hostname)
	if err != nil {
		cli.PrintError(err.Error(), flagJSON)
		os.Exit(2)
	}

	zoneID, err := cfapi.GetZoneID(ctx, api, zone)
	if err != nil {
		cli.PrintError(err.Error(), flagJSON)
		os.Exit(3)
	}

	existing, err := cfapi.FindRecord(ctx, api, zoneID, hostname, recType)
	if err != nil {
		cli.PrintError(err.Error(), flagJSON)
		os.Exit(3)
	}

	if existing == nil {
		cli.PrintError("record not found: "+hostname, flagJSON)
		os.Exit(1)
	}

	if existing.Content == content {
		cli.PrintResult(cli.Result{Action: "no-change", Hostname: hostname, IP: content, Status: "success"}, flagJSON)
		os.Exit(4)
	}

	if err := cfapi.UpdateRecord(ctx, api, zoneID, existing.ID, recType, hostname, content, flagTTL, flagProxied); err != nil {
		cli.PrintError(err.Error(), flagJSON)
		os.Exit(3)
	}
	cli.PrintResult(cli.Result{Action: "update", Hostname: hostname, IP: content, Status: "success"}, flagJSON)
	return nil
}
