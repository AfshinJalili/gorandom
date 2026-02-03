package commands

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/AfshinJalili/gorandom/internal/articles"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open [url|index]",
	Short: "Open article in browser",
	Run:   runOpen,
}

func init() {
	rootCmd.AddCommand(openCmd)
}

func runOpen(cmd *cobra.Command, args []string) {
	var targetUrl string

	if len(args) == 0 {
		// Pick random unread
		pool := loadArticles(cmd)
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

		if len(unread) == 0 {
			// Fallback to random any
			unread = pool
		}

		article := articles.PickRandom(unread)
		if article != nil {
			targetUrl = article.URL
			cmd.Printf("Opening random article: %s\n", article.Title)
		}
	} else {
		input := args[0]
		if index, err := strconv.Atoi(input); err == nil {
			entries, err := historyStore.GetSortedHistory()
			if err != nil {
				printHistoryLoadError(cmd, err)
				return
			}
			idx := index - 1
			if idx >= 0 && idx < len(entries) {
				targetUrl = entries[idx].URL
			} else {
				printInvalidIndex(cmd, input)
				return
			}
		} else {
			targetUrl = input
		}
	}

	if targetUrl == "" {
		printError(cmd, "Could not determine URL to open.", "pass a URL or a history index")
		return
	}

	if err := openBrowser(targetUrl); err != nil {
		printOpenBrowserError(cmd, err)
		cmd.Printf("URL: %s\n", targetUrl)
	}
}

func openBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}
