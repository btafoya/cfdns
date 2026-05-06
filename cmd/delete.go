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

var deleteCmd = &cobra.Command{
	Use:   "delete <hostname>",
	Short: "Delete a DNS record",
	Args:  cobra.ExactArgs(1),
	RunE:  runDelete,
}

func runDelete(cmd *cobra.Command, args []string) error {
	hostname := args[0]

	if err := util.ValidateHostname(hostname); err != nil {
		cli.PrintError(err.Error(), flagJSON)
		os.Exit(2)
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

	existing, err := cfapi.FindRecord(ctx, api, zoneID, hostname, flagType)
	if err != nil {
		cli.PrintError(err.Error(), flagJSON)
		os.Exit(3)
	}

	if existing == nil {
		cli.PrintError("record not found: "+hostname, flagJSON)
		os.Exit(1)
	}

	if !flagForce && !cli.Confirm(fmt.Sprintf("Delete %s %s -> %s? (y/N): ", existing.Type, existing.Name, existing.Content)) {
		fmt.Println("Aborted.")
		os.Exit(0)
	}

	if err := cfapi.DeleteRecord(ctx, api, zoneID, existing.ID); err != nil {
		cli.PrintError(err.Error(), flagJSON)
		os.Exit(3)
	}
	cli.PrintResult(cli.Result{Action: "delete", Hostname: hostname, Status: "success"}, flagJSON)
	return nil
}
