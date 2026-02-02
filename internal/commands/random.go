package commands

import (
	"fmt"
	"io"

	"github.com/AfshinJalili/gorandom/internal/articles"
	"github.com/AfshinJalili/gorandom/internal/output"
	"github.com/AfshinJalili/gorandom/internal/ui"
	"github.com/spf13/cobra"
)

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get one random article",
	Run:   runRandom,
}

func init() {
	rootCmd.AddCommand(randomCmd)
	randomCmd.Flags().BoolP("any", "a", false, "Include already-read articles in random pick")
	randomCmd.Flags().StringP("source", "s", "", "Filter by source")
	randomCmd.Flags().Bool("json", false, "Output JSON instead of interactive UI")
	randomCmd.Flags().Bool("plain", false, "Output plain text instead of interactive UI")
	randomCmd.Flags().Bool("no-ui", false, "Disable interactive UI (same as --plain)")
}

func runRandom(cmd *cobra.Command, args []string) {
	mode, err := resolveOutputMode(cmd)
	if err != nil {
		printOutputModeError(cmd, err)
		return
	}

	any, _ := cmd.Flags().GetBool("any")
	sourceStr, _ := cmd.Flags().GetString("source")

	pool := articles.Data

	pool, err = filterPool(pool, sourceStr, "")
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

		var unread []articles.Article
		for _, a := range pool {
			if !readUrls[a.URL] {
				unread = append(unread, a)
			}
		}

		// If unread is empty, decide whether to fallback or warn
		// In interactive mode, maybe we just show the message or specific screen?
		// For now, let's just warn and exit like before, or fallback (logic from before)
		if len(unread) == 0 {
			msg := "Congratulations! You have read all articles"
			if sourceStr != "" {
				msg += fmt.Sprintf(" from %s", articles.FormatSource(articles.Source(sourceStr)))
			}
			msg += "!"
			cmd.Println(msg)
			cmd.Println("Use --any to get a random article from all (including read).")
			return
		}
		pool = unread
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

func writeArticlePlain(w io.Writer, a *articles.Article) {
	fmt.Fprint(w, output.FormatArticlePlain(toOutputArticle(a)))
}

func writeArticleJSON(w io.Writer, a *articles.Article) error {
	payload, err := output.FormatArticleJSON(toOutputArticle(a))
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, payload)
	return err
}
