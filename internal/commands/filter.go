package commands

import (
	"strings"

	"github.com/AfshinJalili/gorandom/internal/articles"
)

func filterPool(pool []articles.Article, sourceStr string, keyword string) ([]articles.Article, error) {
	if sourceStr != "" {
		if !articles.IsValidSource(sourceStr) {
			return nil, errInvalidSource(sourceStr)
		}
		pool = articles.FilterBySource(pool, articles.Source(sourceStr))
	}

	if keyword == "" {
		return pool, nil
	}

	kw := strings.ToLower(keyword)
	var filtered []articles.Article
	for _, a := range pool {
		if strings.Contains(strings.ToLower(a.Title), kw) ||
			strings.Contains(strings.ToLower(string(a.Source)), kw) {
			filtered = append(filtered, a)
		}
	}
	return filtered, nil
}

type invalidSourceErr struct {
	source string
}

func (e invalidSourceErr) Error() string {
	return "Invalid source: " + e.source
}

func errInvalidSource(source string) error {
	return invalidSourceErr{source: source}
}
