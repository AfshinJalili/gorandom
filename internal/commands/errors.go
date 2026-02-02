package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func printError(cmd *cobra.Command, summary string, hint string) {
	cmd.Println(summary)
	if hint != "" {
		cmd.Printf("Try: %s\n", hint)
	}
}

func printHistoryLoadError(cmd *cobra.Command, err error) {
	printError(cmd, fmt.Sprintf("Could not load history: %v", err), "set GORANDOM_CONFIG_DIR to a writable folder")
}

func printHistorySaveError(cmd *cobra.Command, err error) {
	printError(cmd, fmt.Sprintf("Could not save history: %v", err), "set GORANDOM_CONFIG_DIR to a writable folder")
}

func printInvalidSource(cmd *cobra.Command, source string) {
	printError(cmd, fmt.Sprintf("Invalid source: %s", source), "run \"gorandom sources\"")
}

func printInvalidIndex(cmd *cobra.Command, input string) {
	printError(cmd, fmt.Sprintf("Invalid index: %s", input), "run \"gorandom history\" and use 1 for most recent")
}

func printOutputModeError(cmd *cobra.Command, err error) {
	printError(cmd, err.Error(), "use only one of --json or --plain/--no-ui")
}

func printOpenBrowserError(cmd *cobra.Command, err error) {
	printError(cmd, fmt.Sprintf("Could not open browser: %v", err), "copy and open the URL manually")
}
