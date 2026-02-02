package output

import (
	"strings"
	"testing"
	"time"
)

func TestFormatArticlePlain(t *testing.T) {
	a := Article{Title: "Title", URL: "http://example.com", Source: "docs"}
	out := FormatArticlePlain(a)
	if !strings.Contains(out, "Title") || !strings.Contains(out, "http://example.com") {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestFormatArticleJSON(t *testing.T) {
	a := Article{Title: "Title", URL: "http://example.com", Source: "docs"}
	out, err := FormatArticleJSON(a)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "\"url\"") || !strings.Contains(out, "\"source\"") {
		t.Fatalf("unexpected JSON output: %q", out)
	}
}

func TestFormatHistoryPlain(t *testing.T) {
	entries := []HistoryEntry{
		{
			Title:        "One",
			URL:          "http://example.com/1",
			ViewedAt:     time.Date(2025, 1, 2, 3, 4, 5, 0, time.UTC),
			IsRead:       true,
			IsBookmarked: true,
		},
	}
	out := FormatHistoryPlain(entries)
	if !strings.Contains(out, "read") || !strings.Contains(out, "★") {
		t.Fatalf("unexpected output: %q", out)
	}
	if !strings.Contains(out, "2025-01-02") || !strings.Contains(out, "http://example.com/1") {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestFormatHistoryJSON(t *testing.T) {
	entries := []HistoryEntry{
		{Title: "One", URL: "http://example.com/1"},
	}
	out, err := FormatHistoryJSON(entries)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "\"url\"") || !strings.Contains(out, "[") {
		t.Fatalf("unexpected JSON output: %q", out)
	}
}

func TestFormatArticlesPlain(t *testing.T) {
	items := []Article{
		{Title: "One", URL: "http://example.com/1", Source: "docs"},
		{Title: "Two", URL: "http://example.com/2", Source: "blog"},
	}
	out := FormatArticlesPlain(items)
	if !strings.Contains(out, "One") || !strings.Contains(out, "Two") {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestFormatArticlesJSON(t *testing.T) {
	items := []Article{
		{Title: "One", URL: "http://example.com/1", Source: "docs"},
	}
	out, err := FormatArticlesJSON(items)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "[") || !strings.Contains(out, "\"url\"") {
		t.Fatalf("unexpected JSON output: %q", out)
	}
}
