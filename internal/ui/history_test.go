package ui

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/AfshinJalili/gorandom/internal/history"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type errStore struct {
	markReadErr       error
	markUnreadErr     error
	toggleBookmarkErr error
}

func (e errStore) AddToHistory(url string) error                 { return nil }
func (e errStore) MarkAsRead(url string) error                   { return e.markReadErr }
func (e errStore) MarkAsUnread(url string) (bool, error)         { return e.markUnreadErr == nil, e.markUnreadErr }
func (e errStore) GetReadUrls() (map[string]bool, error)         { return map[string]bool{}, nil }
func (e errStore) GetSortedHistory() ([]history.HistoryEntry, error) {
	return []history.HistoryEntry{}, nil
}
func (e errStore) CalculateStreak() (int, error) { return 0, nil }
func (e errStore) ToggleBookmark(url string) (bool, error) {
	if e.toggleBookmarkErr != nil {
		return false, e.toggleBookmarkErr
	}
	return true, nil
}
func (e errStore) GetBookmarks() ([]history.HistoryEntry, error) { return []history.HistoryEntry{}, nil }

func TestHistoryModelToggleReadError(t *testing.T) {
	entry := history.HistoryEntry{URL: "http://example.com", ViewedAt: time.Now()}
	items := []list.Item{createListItem(entry)}
	keys := newListKeyMap()
	delegate := newItemDelegate(keys)
	l := list.New(items, delegate, 0, 0)

	m := historyModel{
		list:         l,
		keys:         keys,
		historyStore: errStore{markReadErr: errors.New("boom")},
	}

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("m")})
	m2 := updated.(historyModel)

	if !strings.Contains(m2.msg, "Failed to mark as read") {
		t.Fatalf("Expected error message, got %q", m2.msg)
	}

	item := m2.list.SelectedItem().(HistoryItem)
	if item.Entry.IsRead {
		t.Fatal("Expected entry to remain unread on error")
	}
}

func TestHistoryModelToggleBookmarkError(t *testing.T) {
	entry := history.HistoryEntry{URL: "http://example.com", ViewedAt: time.Now()}
	items := []list.Item{createListItem(entry)}
	keys := newListKeyMap()
	delegate := newItemDelegate(keys)
	l := list.New(items, delegate, 0, 0)

	m := historyModel{
		list:         l,
		keys:         keys,
		historyStore: errStore{toggleBookmarkErr: errors.New("boom")},
	}

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("b")})
	m2 := updated.(historyModel)

	if !strings.Contains(m2.msg, "Failed to toggle bookmark") {
		t.Fatalf("Expected error message, got %q", m2.msg)
	}

	item := m2.list.SelectedItem().(HistoryItem)
	if item.Entry.IsBookmarked {
		t.Fatal("Expected entry to remain unbookmarked on error")
	}
}
