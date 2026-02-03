package commands

import (
	"strings"
	"testing"

	"github.com/AfshinJalili/gorandom/internal/articles"
)

func TestFormatSearchTitleIcons(t *testing.T) {
	a := articles.Article{
		Title:  "Example",
		URL:    "http://example.com",
		Source: articles.SourceDocs,
	}

	title := formatSearchTitle(a, true, true)
	if !strings.Contains(title, "✓") || !strings.Contains(title, "★") {
		t.Fatalf("expected read/bookmark icons, got %q", title)
	}

	title = formatSearchTitle(a, false, false)
	if !strings.Contains(title, "○") {
		t.Fatalf("expected unread icon, got %q", title)
	}
	if strings.Contains(title, "★") {
		t.Fatalf("did not expect bookmark icon, got %q", title)
	}
	if !strings.Contains(title, a.Title) {
		t.Fatalf("expected title present, got %q", title)
	}
	if !strings.Contains(title, articles.FormatSource(a.Source)) {
		t.Fatalf("expected source present, got %q", title)
	}
}
