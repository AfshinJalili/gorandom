package commands

import (
	"github.com/AfshinJalili/gorandom/internal/articles"
	"github.com/AfshinJalili/gorandom/internal/history"
	"github.com/AfshinJalili/gorandom/internal/output"
)

func toOutputArticle(a *articles.Article) output.Article {
	if a == nil {
		return output.Article{}
	}
	return output.Article{
		Title:  a.Title,
		URL:    a.URL,
		Source: string(a.Source),
	}
}

func toOutputHistory(entries []history.HistoryEntry) []output.HistoryEntry {
	out := make([]output.HistoryEntry, 0, len(entries))
	for _, e := range entries {
		out = append(out, output.HistoryEntry{
			Title:        findTitle(e.URL),
			URL:          e.URL,
			ViewedAt:     e.ViewedAt,
			IsRead:       e.IsRead,
			IsBookmarked: e.IsBookmarked,
		})
	}
	return out
}

func toOutputArticles(items []articles.Article) []output.Article {
	out := make([]output.Article, 0, len(items))
	for i := range items {
		out = append(out, toOutputArticle(&items[i]))
	}
	return out
}
