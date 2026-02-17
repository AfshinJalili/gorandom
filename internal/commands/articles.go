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
	notifyFetchingIfNeeded(cmd)
	data, err := articles.GetData()
	if err != nil || len(data) == 0 {
		printError(cmd, fmt.Sprintf("Could not load sources: %v", err), "check your network or run \"gorandom sources update\"")
		return nil, false
	}
	return data, true
}

func loadArticlesOptional(cmd *cobra.Command) []articles.Article {
	notifyFetchingIfNeeded(cmd)
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

func notifyFetchingIfNeeded(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	exists, err := articles.CacheExists()
	if err != nil {
		printWarning(cmd, fmt.Sprintf("Could not check sources cache: %v", err))
		return
	}
	if exists {
		return
	}

	sourcesFetchOnce.Do(func() {
		// TUI mode shows a small spinner while fetching on first run.
		if isLikelyTUI(cmd) {
			if v := os.Getenv("GORANDOM_SOURCES_SPINNER"); v == "0" {
				cmd.Println("Fetching sources...")
				return
			}
			ui.ShowSpinner("Fetching sources...", 600*time.Millisecond)
			return
		}
		cmd.Println("Fetching sources...")
	})
}

func isLikelyTUI(cmd *cobra.Command) bool {
	return cmd.OutOrStdout() == os.Stdout
}
