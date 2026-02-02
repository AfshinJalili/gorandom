package history

import (
	"os"
	"testing"
	"time"
)

func TestExtensions(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gorandom-ext-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	os.Setenv("GORANDOM_CONFIG_DIR", tmpDir)
	defer os.Unsetenv("GORANDOM_CONFIG_DIR")

	url := "http://example.com/streak"

	// Test Bookmark
	booked, err := ToggleBookmark(url)
	if err != nil {
		t.Fatal(err)
	}
	if !booked {
		t.Error("Expected booked=true")
	}

	bookmarks, err := GetBookmarks()
	if err != nil {
		t.Fatal(err)
	}
	if len(bookmarks) != 1 || bookmarks[0].URL != url {
		t.Error("Expected 1 bookmark")
	}

	// Toggle off
	booked, err = ToggleBookmark(url)
	if err != nil {
		t.Fatal(err)
	}
	if booked {
		t.Error("Expected booked=false")
	}

	// Test Streak
	// Mark as read today
	if err := MarkAsRead(url); err != nil {
		t.Fatal(err)
	}

	streak, err := CalculateStreak()
	if err != nil {
		t.Fatal(err)
	}
	if streak != 1 {
		t.Errorf("Expected streak 1, got %d", streak)
	}

	// Manipulate history to simulate yesterday
	history, _ := LoadHistory()
	yesterday := time.Now().AddDate(0, 0, -1)
	history.Entries[0].ReadAt = yesterday // current entry moved to yesterday
	SaveHistory(history)

	// Add another for today
	url2 := "http://example.com/today"
	MarkAsRead(url2)

	streak, err = CalculateStreak()
	if err != nil {
		t.Fatal(err)
	}
	if streak != 2 {
		t.Errorf("Expected streak 2, got %d", streak)
	}
}
