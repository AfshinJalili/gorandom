package commands

import (
	"fmt"
	"strconv"

	"github.com/AfshinJalili/gorandom/internal/ui"
	"github.com/spf13/cobra"
)

var bookmarkCmd = &cobra.Command{
	Use:   "bookmark [url|index]",
	Short: "Toggle bookmark for an article",
	Run:   runBookmark,
}

var bookmarksCmd = &cobra.Command{
	Use:   "bookmarks",
	Short: "List bookmarked articles",
	Run:   runListBookmarks,
}

func init() {
	rootCmd.AddCommand(bookmarkCmd)
	rootCmd.AddCommand(bookmarksCmd)
}

func runBookmark(cmd *cobra.Command, args []string) {
	var targetUrl string

	if len(args) == 0 {
		// Interactive mode: Open history list
		entries, err := historyStore.GetSortedHistory()
		if err != nil {
			printHistoryLoadError(cmd, err)
			return
		}
		if _, err := ui.ShowHistory(entries, "Toggle Bookmarks (press 'b')"); err != nil {
			printError(cmd, fmt.Sprintf("Could not show history: %v", err), "pass a URL or history index")
		}
		return
	} else {
		input := args[0]
		if index, err := strconv.Atoi(input); err == nil {
			entries, _ := historyStore.GetSortedHistory()
			idx := index - 1
			if idx >= 0 && idx < len(entries) {
				targetUrl = entries[idx].URL
			}
		} else {
			targetUrl = input
		}
	}

	if targetUrl == "" {
		printError(cmd, "Invalid URL or index.", "pass a valid URL or a history index")
		return
	}

	bookmarked, err := historyStore.ToggleBookmark(targetUrl)
	if err != nil {
		printHistorySaveError(cmd, err)
		return
	}

	status := "Bookmarked"
	if !bookmarked {
		status = "Removed bookmark"
	}

	title := findTitle(targetUrl)
	cmd.Printf("%s: %s\n", status, title)
}

func runListBookmarks(cmd *cobra.Command, args []string) {
	bookmarks, err := historyStore.GetBookmarks()
	if err != nil {
		printHistoryLoadError(cmd, err)
		return
	}

	if len(bookmarks) == 0 {
		cmd.Println("No bookmarks yet.")
		return
	}

	// User interacts within the UI
	if _, err := ui.ShowHistory(bookmarks, "Bookmarks"); err != nil {
		printError(cmd, fmt.Sprintf("Could not show bookmarks: %v", err), "re-run without UI using --plain")
	}
}
