package commands

import (
	"github.com/AfshinJalili/gorandom/internal/articles"
	"github.com/spf13/cobra"
)

var sourcesCmd = &cobra.Command{
	Use:   "sources",
	Short: "List content sources",
	Run:   runSources,
}

func init() {
	rootCmd.AddCommand(sourcesCmd)
}

func runSources(cmd *cobra.Command, args []string) {
	cmd.Println("\nAvailable sources:")
	for _, source := range articles.Sources {
		count := len(articles.FilterBySource(articles.Data, source))
		// Format similar to original CLI
		// using padding manually since no nice tabulated print
		name := string(source)
		desc := articles.FormatSource(source)
		cmd.Printf("  %-15s - %s (%d articles)\n", name, desc, count)
	}
	cmd.Println()
}
