package commands

import (
	"fmt"
	"io"

	"github.com/AfshinJalili/gorandom/internal/articles"
	"github.com/AfshinJalili/gorandom/internal/output"
	"github.com/AfshinJalili/gorandom/internal/ui"
	"github.com/spf13/cobra"
)

var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Get the next unread article",
	Run:   runNext,
}

func init() {
	rootCmd.AddCommand(nextCmd)
	nextCmd.Flags().BoolP("any", "a", false, "Include already-read articles in pick")
	nextCmd.Flags().StringP("source", "s", "", "Filter by source")
	nextCmd.Flags().Bool("json", false, "Output JSON instead of interactive UI")
	nextCmd.Flags().Bool("plain", false, "Output plain text instead of interactive UI")
	nextCmd.Flags().Bool("no-ui", false, "Disable interactive UI (same as --plain)")
}

func runNext(cmd *cobra.Command, args []string) {
	mode, err := resolveOutputMode(cmd)
	if err != nil {
		printOutputModeError(cmd, err)
		return
	}

	any, _ := cmd.Flags().GetBool("any")
	sourceStr, _ := cmd.Flags().GetString("source")

	pool, err := filterPool(articles.Data, sourceStr, "")
	if err != nil {
		printInvalidSource(cmd, sourceStr)
		return
	}
	if len(pool) == 0 {
		cmd.Printf("No articles found for source: %s\n", sourceStr)
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
			msg := "Congratulations! You have read all articles"
			if sourceStr != "" {
				msg += fmt.Sprintf(" from %s", articles.FormatSource(articles.Source(sourceStr)))
			}
			msg += "!"
			cmd.Println(msg)
			cmd.Println("Use --any to get a random article from all (including read).")
			return
		}
	}

	article := articles.PickRandom(pool)
	if article == nil {
		cmd.Println("No articles available.")
		return
	}

	switch mode {
	case outputUI:
		if err := ui.ShowRandomArticle(article, pool); err != nil {
			printError(cmd, fmt.Sprintf("Could not show article: %v", err), "re-run without UI using --plain")
		}
	case outputPlain:
		if err := historyStore.AddToHistory(article.URL); err != nil {
			printHistorySaveError(cmd, err)
			return
		}
		writeArticlePlain(cmd.OutOrStdout(), article)
	case outputJSON:
		if err := historyStore.AddToHistory(article.URL); err != nil {
			printHistorySaveError(cmd, err)
			return
		}
		if err := writeArticleJSON(cmd.OutOrStdout(), article); err != nil {
			printError(cmd, fmt.Sprintf("Could not write JSON: %v", err), "re-run with --plain")
		}
	}
}

func writeArticlesPlain(w io.Writer, items []articles.Article) {
	payload := output.FormatArticlesPlain(toOutputArticles(items))
	io.WriteString(w, payload)
}

func writeArticlesJSON(w io.Writer, items []articles.Article) error {
	payload, err := output.FormatArticlesJSON(toOutputArticles(items))
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, payload)
	return err
}
