package commands

import (
	"fmt"
	"time"

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
	sourcesCmd.AddCommand(sourcesUpdateCmd)
	sourcesCmd.AddCommand(sourcesStatusCmd)
}

func runSources(cmd *cobra.Command, args []string) {
	cmd.Println("\nAvailable sources:")
	data, ok := loadArticlesRequired(cmd)
	if !ok {
		return
	}
	for _, source := range articles.Sources {
		count := len(articles.FilterBySource(data, source))
		// Format similar to original CLI
		// using padding manually since no nice tabulated print
		name := string(source)
		desc := articles.FormatSource(source)
		cmd.Printf("  %-15s - %s (%d articles)\n", name, desc, count)
	}
	cmd.Println()
}

var sourcesUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Fetch latest sources and refresh cache",
	Run:   runSourcesUpdate,
}

var sourcesStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show sources cache status",
	Run:   runSourcesStatus,
}

func runSourcesUpdate(cmd *cobra.Command, args []string) {
	updated, err := articles.ForceUpdate()
	if err != nil {
		printError(cmd, fmt.Sprintf("Could not update sources: %v", err), "check your network or GORANDOM_SOURCES_URL")
		return
	}
	if updated {
		cmd.Println("Sources updated.")
		return
	}
	cmd.Println("Sources are already up to date.")
}

func runSourcesStatus(cmd *cobra.Command, args []string) {
	status, err := articles.CacheStatusInfo()
	if err != nil {
		printError(cmd, fmt.Sprintf("Could not read cache status: %v", err), "check your config directory")
		return
	}
	cmd.Printf("URL: %s\n", status.URL)
	if status.FetchedAt.IsZero() {
		cmd.Println("Cache: none")
		return
	}
	cmd.Printf("Cache: %s\n", status.CachePath)
	cmd.Printf("Fetched: %s\n", status.FetchedAt.Format(time.RFC3339))
	cmd.Printf("Age: %s\n", status.Age.Round(time.Second))
	cmd.Printf("Stale: %t\n", status.Stale)
}
