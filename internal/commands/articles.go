package commands

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/AfshinJalili/gorandom/internal/articles"
	"github.com/AfshinJalili/gorandom/internal/ui"
	"github.com/spf13/cobra"
)

var sourcesFetchOnce sync.Once

func loadArticlesRequired(cmd *cobra.Command) ([]articles.Article, bool) {
	maybeNotifyFetchingSpinner(cmd)
	maybeNotifyFetching(cmd)
	data, err := articles.GetData()
	if err != nil || len(data) == 0 {
		printError(cmd, fmt.Sprintf("Could not load sources: %v", err), "check your network or run \"gorandom sources update\"")
		return nil, false
	}
	return data, true
}

func loadArticlesOptional(cmd *cobra.Command) []articles.Article {
	maybeNotifyFetchingSpinner(cmd)
	maybeNotifyFetching(cmd)
	data, err := articles.GetData()
	if err != nil || len(data) == 0 {
		printWarning(cmd, "Sources cache missing; run \"gorandom sources update\".")
		return nil
	}
	return data
}

func loadArticlesQuiet() []articles.Article {
	data, err := articles.GetData()
	if err != nil {
		return nil
	}
	return data
}

func maybeNotifyFetching(cmd *cobra.Command) {
	exists, err := articles.CacheExists()
	if err != nil {
		printWarning(cmd, fmt.Sprintf("Could not check sources cache: %v", err))
	}
	if !exists {
		sourcesFetchOnce.Do(func() {
			cmd.Println("Fetching sources...")
		})
	}
}

func maybeNotifyFetchingSpinner(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	if cmd.OutOrStdout() == nil {
		return
	}
	// TUI mode shows a small spinner while fetching on first run.
	if isLikelyTUI(cmd) {
		if v := os.Getenv("GORANDOM_SOURCES_SPINNER"); v == "0" {
			return
		}
		sourcesFetchOnce.Do(func() {
			ui.ShowSpinner("Fetching sources...", 600*time.Millisecond)
		})
	}
}

func isLikelyTUI(cmd *cobra.Command) bool {
	return cmd.OutOrStdout() == os.Stdout
}
