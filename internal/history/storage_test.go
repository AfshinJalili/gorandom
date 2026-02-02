package history

import (
	"os"
	"testing"
	"time"
)

func TestStorage(t *testing.T) {
	// Setup temp config dir
	tmpDir, err := os.MkdirTemp("", "gorandom-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	os.Setenv("GORANDOM_CONFIG_DIR", tmpDir)
	defer os.Unsetenv("GORANDOM_CONFIG_DIR")

	// Test LoadHistory (empty)
	history, err := LoadHistory()
	if err != nil {
		t.Fatalf("LoadHistory failed: %v", err)
	}
	if len(history.Entries) != 0 {
		t.Error("Expected empty history")
	}

	// Test AddToHistory
	url1 := "http://example.com/1"
	if err := AddToHistory(url1); err != nil {
		t.Fatalf("AddToHistory failed: %v", err)
	}

	history, err = LoadHistory()
	if err != nil {
		t.Fatal(err)
	}
	if len(history.Entries) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(history.Entries))
	}
	if history.Entries[0].URL != url1 {
		t.Errorf("Expected URL %s, got %s", url1, history.Entries[0].URL)
	}
	if history.Entries[0].IsRead {
		t.Error("Expected IsRead to be false")
	}

	// Test MarkAsRead
	if err := MarkAsRead(url1); err != nil {
		t.Fatalf("MarkAsRead failed: %v", err)
	}

	readUrls, err := GetReadUrls()
	if err != nil {
		t.Fatal(err)
	}
	if !readUrls[url1] {
		t.Error("Expected URL to be marked read")
	}

	// Test MarkAsUnread
	success, err := MarkAsUnread(url1)
	if err != nil {
		t.Fatal(err)
	}
	if !success {
		t.Error("Expected MarkAsUnread to return success=true")
	}

	readUrls, err = GetReadUrls()
	if err != nil {
		t.Fatal(err)
	}
	if readUrls[url1] {
		t.Error("Expected URL to be marked unread")
	}
}

func TestGetSortedHistory(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gorandom-sorted-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	os.Setenv("GORANDOM_CONFIG_DIR", tmpDir)
	defer os.Unsetenv("GORANDOM_CONFIG_DIR")

	url1 := "http://example.com/old"
	url2 := "http://example.com/new"

	if err := AddToHistory(url1); err != nil {
		t.Fatal(err)
	}
	// Sleep briefly to ensure timestamp difference
	time.Sleep(10 * time.Millisecond)
	if err := AddToHistory(url2); err != nil {
		t.Fatal(err)
	}

	sorted, err := GetSortedHistory()
	if err != nil {
		t.Fatal(err)
	}
	if len(sorted) != 2 {
		t.Fatalf("Expected 2 entries, got %d", len(sorted))
	}

	// Newest should be first
	if sorted[0].URL != url2 {
		t.Errorf("Expected newest entry (url2) first, got %s", sorted[0].URL)
	}
}

func TestMarkAsReadNewEntry(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gorandom-mark-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	os.Setenv("GORANDOM_CONFIG_DIR", tmpDir)
	defer os.Unsetenv("GORANDOM_CONFIG_DIR")

	// Mark a URL as read that doesn't exist yet
	url := "http://new.example.com"
	err = MarkAsRead(url)
	if err != nil {
		t.Fatal(err)
	}

	h, _ := LoadHistory()
	if len(h.Entries) != 1 {
		t.Fatalf("Expected 1 entry, got %d", len(h.Entries))
	}
	if !h.Entries[0].IsRead {
		t.Error("Expected new entry to be created and marked read")
	}
	if h.Entries[0].ReadAt.IsZero() {
		t.Error("Expected ReadAt to be set")
	}
}

func TestLoadHistoryCorruptJSON(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gorandom-corrupt-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	os.Setenv("GORANDOM_CONFIG_DIR", tmpDir)
	defer os.Unsetenv("GORANDOM_CONFIG_DIR")

	path, err := GetHistoryFilePath()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		t.Fatal(err)
	}

	bad := []byte("{not valid json")
	if err := os.WriteFile(path, bad, 0644); err != nil {
		t.Fatal(err)
	}

	if _, err := LoadHistory(); err == nil {
		t.Fatal("Expected LoadHistory to fail for corrupted JSON")
	}

	if err := AddToHistory("http://example.com/corrupt"); err == nil {
		t.Fatal("Expected AddToHistory to fail when history is corrupted")
	}

	after, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(after) != string(bad) {
		t.Error("Expected corrupted history file to remain unchanged on error")
	}
}
