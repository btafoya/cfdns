package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/btafoya/cfdns/internal/cli"
	cfapi "github.com/btafoya/cfdns/internal/cloudflare"
	"github.com/btafoya/cfdns/internal/util"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <hostname> <content>",
	Short: "Create a DNS record",
	Args:  cobra.ExactArgs(2),
	RunE:  runCreate,
}

func runCreate(cmd *cobra.Command, args []string) error {
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

	if existing != nil {
		if existing.Content == content {
			cli.PrintResult(cli.Result{Action: "no-change", Hostname: hostname, IP: content, Status: "success"}, flagJSON)
			os.Exit(4)
		}

		fmt.Printf("Record already exists: %s -> %s\n", existing.Name, existing.Content)

		if !flagForce && !cli.Confirm("Update record with new IP? (y/N): ") {
			fmt.Println("Aborted.")
			os.Exit(0)
		}

		if err := cfapi.UpdateRecord(ctx, api, zoneID, existing.ID, recType, hostname, content, flagTTL, flagProxied); err != nil {
			cli.PrintError(err.Error(), flagJSON)
			os.Exit(3)
		}
		cli.PrintResult(cli.Result{Action: "update", Hostname: hostname, IP: content, Status: "success"}, flagJSON)
		return nil
	}

	if err := cfapi.CreateRecord(ctx, api, zoneID, recType, hostname, content, flagTTL, flagProxied); err != nil {
		cli.PrintError(err.Error(), flagJSON)
		os.Exit(3)
	}
	cli.PrintResult(cli.Result{Action: "create", Hostname: hostname, IP: content, Status: "success"}, flagJSON)
	return nil
}

func resolveType(content string) string {
	if flagType != "A" {
		return flagType
	}
	if util.IsIPv6(content) {
		return "AAAA"
	}
	return "A"
}
