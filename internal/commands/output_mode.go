package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

type outputMode int

const (
	outputUI outputMode = iota
	outputPlain
	outputJSON
)

func resolveOutputMode(cmd *cobra.Command) (outputMode, error) {
	jsonOut, _ := cmd.Flags().GetBool("json")
	plainOut, _ := cmd.Flags().GetBool("plain")
	noUI, _ := cmd.Flags().GetBool("no-ui")

	if jsonOut && (plainOut || noUI) {
		return outputUI, fmt.Errorf("use only one of --json or --plain/--no-ui")
	}
	if jsonOut {
		return outputJSON, nil
	}
	if plainOut || noUI {
		return outputPlain, nil
	}
	return outputUI, nil
}
