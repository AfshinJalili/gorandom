package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AfshinJalili/gorandom/internal/history"
	"github.com/AfshinJalili/gorandom/internal/ui"
	"github.com/spf13/cobra"
)

var markCmd = &cobra.Command{
	Use:   "mark [url|index]",
	Short: "Mark an article as read",
	Long:  "Mark an article as read. Pass URL, history index (1 = most recent), or no arg for interactive pick.",
	Run: func(cmd *cobra.Command, args []string) {
		runMarkUnmark(cmd, args, true)
	},
}

var unmarkCmd = &cobra.Command{
	Use:   "unmark [url|index]",
	Short: "Mark an article as unread",
	Long:  "Mark an article as unread. Pass URL, history index (1 = most recent), or no arg for interactive pick.",
	Run: func(cmd *cobra.Command, args []string) {
		runMarkUnmark(cmd, args, false)
	},
}

func init() {
	rootCmd.AddCommand(markCmd)
	rootCmd.AddCommand(unmarkCmd)
}

func runMarkUnmark(cmd *cobra.Command, args []string, markRead bool) {
	var targetUrl string

	if len(args) == 0 {
		// Interactive mode
		loadArticles(cmd)
		entries, err := historyStore.GetSortedHistory()
		if err != nil {
			printHistoryLoadError(cmd, err)
			return
		}

		var filtered []history.HistoryEntry
		for _, entry := range entries {
			// markRead (mark as read) -> show unread (!entry.IsRead)
			// !markRead (unmark/mark as unread) -> show read (entry.IsRead)
			if (markRead && !entry.IsRead) || (!markRead && entry.IsRead) {
				filtered = append(filtered, entry)
			}
		}

		if len(filtered) == 0 {
			if markRead {
				cmd.Println("No unread articles to mark.")
			} else {
				cmd.Println("No read articles to unmark.")
			}
			return
		}

		// Show interactive list
		title := "Mark as Read (press 'm')"
		if !markRead {
			title = "Mark as Unread (press 'm')"
		}

		if _, err := ui.ShowHistory(filtered, title); err != nil {
			printError(cmd, fmt.Sprintf("Could not show history: %v", err), "pass a URL or history index")
		}
		return

	} else {
		// Args provided
		input := args[0]
		if index, err := strconv.Atoi(input); err == nil {
			// It's a number (index)
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
				os.Exit(1)
			}
		} else {
			// It's a URL
			targetUrl = input
		}
	}

	var err error
	if markRead {
		err = historyStore.MarkAsRead(targetUrl)
		if err == nil {
			title := findTitle(targetUrl)
			cmd.Printf("Marked as read: %s\n", title)
		}
	} else {
		var success bool
		success, err = historyStore.MarkAsUnread(targetUrl)
		if err == nil {
			if success {
				title := findTitle(targetUrl)
				cmd.Printf("Marked as unread: %s\n", title)
			} else {
				cmd.Printf("URL not found in history: %s\n", targetUrl)
			}
		}
	}

	if err != nil {
		printHistorySaveError(cmd, err)
	}
}

func findTitle(url string) string {
	for _, a := range loadArticlesQuiet() {
		if a.URL == url {
			if a.Title != "" {
				return a.Title
			}
			break
		}
	}
	return url
}
