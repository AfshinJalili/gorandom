package ui

import (
	"strings"
	"testing"
)

func TestListLegend(t *testing.T) {
	legend := listLegend()
	want := []string{"✓ read", "○ unread", "★ bookmarked"}
	for _, token := range want {
		if !strings.Contains(legend, token) {
			t.Fatalf("expected legend to contain %q, got %q", token, legend)
		}
	}
	if !strings.Contains(legend, "•") {
		t.Fatalf("expected legend to use separators, got %q", legend)
	}
}
