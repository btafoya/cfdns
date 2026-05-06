package cli

import (
	"encoding/json"
	"fmt"
	"os"
)

type Result struct {
	Action   string `json:"action"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip,omitempty"`
	Status   string `json:"status"`
	Message  string `json:"message,omitempty"`
}

func PrintResult(r Result, jsonMode bool) {
	if jsonMode {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		_ = enc.Encode(r)
		return
	}
	switch r.Action {
	case "create":
		fmt.Printf("Created: %s -> %s\n", r.Hostname, r.IP)
	case "update":
		fmt.Printf("Updated: %s -> %s\n", r.Hostname, r.IP)
	case "delete":
		fmt.Printf("Deleted: %s\n", r.Hostname)
	case "list":
		if r.Message != "" {
			fmt.Println(r.Message)
		}
	case "no-change":
		fmt.Printf("No change needed: %s already points to %s\n", r.Hostname, r.IP)
	default:
		if r.Message != "" {
			fmt.Println(r.Message)
		}
	}
}

func PrintError(msg string, jsonMode bool) {
	if jsonMode {
		enc := json.NewEncoder(os.Stderr)
		_ = enc.Encode(map[string]string{"status": "error", "message": msg})
		return
	}
	fmt.Fprintln(os.Stderr, "Error:", msg)
}
