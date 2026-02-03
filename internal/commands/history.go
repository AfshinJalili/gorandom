package commands

import (
	"fmt"
	"io"

	"github.com/AfshinJalili/gorandom/internal/history"
	"github.com/AfshinJalili/gorandom/internal/output"
	"github.com/AfshinJalili/gorandom/internal/ui"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "List recently viewed articles",
	Run:   runHistory,
}

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.Flags().Int("limit", 10, "Max entries to show")
	historyCmd.Flags().Bool("json", false, "Output JSON instead of interactive UI")
	historyCmd.Flags().Bool("plain", false, "Output plain text instead of interactive UI")
	historyCmd.Flags().Bool("no-ui", false, "Disable interactive UI (same as --plain)")
}

func runHistory(cmd *cobra.Command, args []string) {
	mode, err := resolveOutputMode(cmd)
	if err != nil {
		printOutputModeError(cmd, err)
		return
	}

	limit, _ := cmd.Flags().GetInt("limit")
	loadArticlesOptional(cmd)

	entries, err := historyStore.GetSortedHistory()
	if err != nil {
		printHistoryLoadError(cmd, err)
		return
	}

	if len(entries) == 0 {
		cmd.Println("No history yet. Run \"gorandom\" to get your first article!")
		return
	}

	toShow := entries
	if len(toShow) > limit {
		toShow = toShow[:limit]
	}

	switch mode {
	case outputUI:
		if _, err := ui.ShowHistory(toShow, "Recent History"); err != nil {
			printError(cmd, fmt.Sprintf("Could not show history: %v", err), "re-run without UI using --plain")
		}
	case outputPlain:
		writeHistoryPlain(cmd.OutOrStdout(), toShow)
	case outputJSON:
		if err := writeHistoryJSON(cmd.OutOrStdout(), toShow); err != nil {
			printError(cmd, fmt.Sprintf("Could not write JSON: %v", err), "re-run with --plain")
		}
	}
}

func writeHistoryPlain(w io.Writer, entries []history.HistoryEntry) {
	payload := output.FormatHistoryPlain(toOutputHistory(entries))
	io.WriteString(w, payload)
}

func writeHistoryJSON(w io.Writer, entries []history.HistoryEntry) error {
	payload, err := output.FormatHistoryJSON(toOutputHistory(entries))
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, payload)
	return err
}
