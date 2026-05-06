package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	flagJSON    bool
	flagForce   bool
	flagType    string
	flagTTL     int
	flagProxied bool
)

var rootCmd = &cobra.Command{
	Use:   "cfdns",
	Short: "Cloudflare DNS management CLI",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&flagJSON, "json", false, "JSON output")
	rootCmd.PersistentFlags().BoolVar(&flagForce, "force", false, "skip confirmation prompts")
	rootCmd.PersistentFlags().StringVar(&flagType, "type", "A", "record type (A, AAAA, CNAME, TXT)")
	rootCmd.PersistentFlags().IntVar(&flagTTL, "ttl", 1, "TTL (1 = auto)")
	rootCmd.PersistentFlags().BoolVar(&flagProxied, "proxied", false, "enable Cloudflare proxy")

	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(listCmd)
}
