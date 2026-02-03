package ui

import (
	"strings"
	"testing"

	"github.com/charmbracelet/bubbles/list"
)

func TestSelectArticleViewShowsLegend(t *testing.T) {
	items := []struct{ Title, Value string }{
		{Title: "Item", Value: "http://example.com"},
	}

	var listItems []list.Item
	for _, i := range items {
		listItems = append(listItems, item{title: i.Title, value: i.Value, desc: i.Value})
	}

	l := list.New(listItems, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Test"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = TitleStyle
	l.SetHeight(20)

	m := model{list: l}
	view := m.View()

	if !strings.Contains(view, "✓ read") || !strings.Contains(view, "○ unread") || !strings.Contains(view, "★ bookmarked") {
		t.Fatalf("expected legend in view, got %q", view)
	}
}
