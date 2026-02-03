package commands

import (
	"fmt"
	"strings"

	"github.com/AfshinJalili/gorandom/internal/articles"
	"github.com/AfshinJalili/gorandom/internal/ui"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search <keyword>",
	Short: "Search articles by title or source",
	Run:   runSearch,
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolP("any", "a", false, "Include already-read articles in results")
	searchCmd.Flags().StringP("source", "s", "", "Filter by source")
	searchCmd.Flags().Bool("json", false, "Output JSON instead of interactive UI")
	searchCmd.Flags().Bool("plain", false, "Output plain text instead of interactive UI")
	searchCmd.Flags().Bool("no-ui", false, "Disable interactive UI (same as --plain)")
}

func runSearch(cmd *cobra.Command, args []string) {
	mode, err := resolveOutputMode(cmd)
	if err != nil {
		printOutputModeError(cmd, err)
		return
	}

	if len(args) == 0 {
		printError(cmd, "Missing keyword.", "run \"gorandom search <keyword>\"")
		return
	}

	keyword := strings.Join(args, " ")
	any, _ := cmd.Flags().GetBool("any")
	sourceStr, _ := cmd.Flags().GetString("source")

	allArticles := loadArticles(cmd)
	pool, err := filterPool(allArticles, sourceStr, keyword)
	if err != nil {
		printInvalidSource(cmd, sourceStr)
		return
	}

	if len(pool) == 0 {
		cmd.Printf("No articles found matching: %s\n", keyword)
		return
	}

	if !any {
		readUrls, err := historyStore.GetReadUrls()
		if err != nil {
			printHistoryLoadError(cmd, err)
			return
		}
		pool = articles.GetUnreadArticles(pool, readUrls)
		if len(pool) == 0 {
			cmd.Printf("No unread articles found matching: %s\n", keyword)
			cmd.Println("Use --any to include read articles.")
			return
		}
	}

	switch mode {
	case outputUI:
		if len(pool) == 1 {
			if err := ui.ShowRandomArticle(&pool[0], pool, len(allArticles)); err != nil {
				printError(cmd, fmt.Sprintf("Could not show article: %v", err), "re-run without UI using --plain")
			}
			return
		}

		var items []struct{ Title, Value string }
		for _, a := range pool {
			title := fmt.Sprintf("%s (%s)", a.Title, articles.FormatSource(a.Source))
			items = append(items, struct{ Title, Value string }{Title: title, Value: a.URL})
		}
		choice, err := ui.SelectArticle(items, "Search Results")
		if err != nil {
			printError(cmd, fmt.Sprintf("Could not show search results: %v", err), "re-run without UI using --plain")
			return
		}
		if choice == "" {
			return
		}
		for i := range pool {
			if pool[i].URL == choice {
				if err := ui.ShowRandomArticle(&pool[i], pool, len(allArticles)); err != nil {
					printError(cmd, fmt.Sprintf("Could not show article: %v", err), "re-run without UI using --plain")
				}
				return
			}
		}
	case outputPlain:
		writeArticlesPlain(cmd.OutOrStdout(), pool)
	case outputJSON:
		if err := writeArticlesJSON(cmd.OutOrStdout(), pool); err != nil {
			printError(cmd, fmt.Sprintf("Could not write JSON: %v", err), "re-run with --plain")
		}
	}
}
