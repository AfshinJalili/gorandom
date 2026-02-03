package ui

import (
	"strings"
	"testing"
	"time"

	"github.com/AfshinJalili/gorandom/internal/articles"
	"github.com/AfshinJalili/gorandom/internal/history"
)

func TestCreateListItemIcons(t *testing.T) {
	orig := articles.Data
	articles.Data = []articles.Article{
		{URL: "http://example.com/1", Title: "One", Source: articles.SourceDocs},
	}
	defer func() { articles.Data = orig }()

	entry := history.HistoryEntry{
		URL:          "http://example.com/1",
		ViewedAt:     time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		IsRead:       true,
		IsBookmarked: true,
	}

	item := createListItem(entry)
	if !strings.Contains(item.TitleStr, "✓") || !strings.Contains(item.TitleStr, "★") {
		t.Fatalf("expected read/bookmark icons in title, got %q", item.TitleStr)
	}

	entry.IsRead = false
	entry.IsBookmarked = false
	item = createListItem(entry)
	if !strings.Contains(item.TitleStr, "○") {
		t.Fatalf("expected unread icon in title, got %q", item.TitleStr)
	}
	if strings.Contains(item.TitleStr, "★") {
		t.Fatalf("did not expect bookmark icon in title, got %q", item.TitleStr)
	}
}
