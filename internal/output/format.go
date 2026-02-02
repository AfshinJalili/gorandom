package output

import (
	"encoding/json"
	"fmt"
	"strings"
)

func FormatArticlePlain(a Article) string {
	return fmt.Sprintf("%s\n%s\n%s\n", a.Title, a.URL, a.Source)
}

func FormatArticleJSON(a Article) (string, error) {
	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data) + "\n", nil
}

func FormatArticlesPlain(items []Article) string {
	var b strings.Builder
	for _, a := range items {
		b.WriteString(FormatArticlePlain(a))
		b.WriteString("\n")
	}
	return b.String()
}

func FormatArticlesJSON(items []Article) (string, error) {
	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data) + "\n", nil
}

func FormatHistoryPlain(entries []HistoryEntry) string {
	var b strings.Builder
	for _, e := range entries {
		read := "unread"
		if e.IsRead {
			read = "read"
		}
		bookmark := " "
		if e.IsBookmarked {
			bookmark = "★"
		}
		date := e.ViewedAt.Format("2006-01-02")
		b.WriteString(fmt.Sprintf("%s %s %s - %s\n", read, bookmark, date, e.Title))
		b.WriteString(fmt.Sprintf("%s\n", e.URL))
	}
	return b.String()
}

func FormatHistoryJSON(entries []HistoryEntry) (string, error) {
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data) + "\n", nil
}
