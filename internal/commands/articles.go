package commands

import (
	"fmt"

	"github.com/AfshinJalili/gorandom/internal/articles"
	"github.com/spf13/cobra"
)

func loadArticles(cmd *cobra.Command) []articles.Article {
	if _, err := articles.SyncIfStale(); err != nil {
		printWarning(cmd, fmt.Sprintf("Could not refresh sources: %v", err))
	}
	data, err := articles.GetData()
	if err != nil {
		printWarning(cmd, fmt.Sprintf("Using bundled sources due to cache error: %v", err))
	}
	return data
}

func loadArticlesQuiet() []articles.Article {
	data, err := articles.GetData()
	if err != nil {
		return articles.Data
	}
	return data
}
