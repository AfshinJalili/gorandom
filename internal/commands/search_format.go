package commands

import (
	"fmt"

	"github.com/AfshinJalili/gorandom/internal/articles"
)

func formatSearchTitle(a articles.Article, isRead bool, isBookmarked bool) string {
	readIcon := "○"
	if isRead {
		readIcon = "✓"
	}
	bookmarkIcon := " "
	if isBookmarked {
		bookmarkIcon = "★"
	}
	return fmt.Sprintf("%s %s %s (%s)", readIcon, bookmarkIcon, a.Title, articles.FormatSource(a.Source))
}
