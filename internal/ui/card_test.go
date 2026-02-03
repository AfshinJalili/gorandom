package ui

import (
	"errors"
	"strings"
	"testing"

	"github.com/AfshinJalili/gorandom/internal/articles"
	"github.com/AfshinJalili/gorandom/internal/history"
	tea "github.com/charmbracelet/bubbletea"
)

type cardStore struct {
	markReadErr       error
	toggleBookmarkErr error
	readUrls          map[string]bool
	readErr           error
	streak            int
	streakErr         error
}

func (s cardStore) AddToHistory(url string) error                     { return nil }
func (s cardStore) MarkAsRead(url string) error                       { return s.markReadErr }
func (s cardStore) MarkAsUnread(url string) (bool, error)             { return true, nil }
func (s cardStore) GetReadUrls() (map[string]bool, error)             { return s.readUrls, s.readErr }
func (s cardStore) GetSortedHistory() ([]history.HistoryEntry, error) { return nil, nil }
func (s cardStore) CalculateStreak() (int, error)                     { return s.streak, s.streakErr }
func (s cardStore) ToggleBookmark(url string) (bool, error) {
	if s.toggleBookmarkErr != nil {
		return false, s.toggleBookmarkErr
	}
	return true, nil
}
func (s cardStore) GetBookmarks() ([]history.HistoryEntry, error) { return nil, nil }

type mutableStore struct {
	read      map[string]bool
	bookmarks map[string]bool
}

func (s *mutableStore) AddToHistory(url string) error         { return nil }
func (s *mutableStore) MarkAsRead(url string) error           { s.read[url] = true; return nil }
func (s *mutableStore) MarkAsUnread(url string) (bool, error) { s.read[url] = false; return true, nil }
func (s *mutableStore) GetReadUrls() (map[string]bool, error) { return s.read, nil }
func (s *mutableStore) GetSortedHistory() ([]history.HistoryEntry, error) {
	return nil, nil
}
func (s *mutableStore) CalculateStreak() (int, error) { return 0, nil }
func (s *mutableStore) ToggleBookmark(url string) (bool, error) {
	if s.bookmarks[url] {
		delete(s.bookmarks, url)
		return false, nil
	}
	s.bookmarks[url] = true
	return true, nil
}
func (s *mutableStore) GetBookmarks() ([]history.HistoryEntry, error) {
	var entries []history.HistoryEntry
	for url := range s.bookmarks {
		entries = append(entries, history.HistoryEntry{URL: url, IsBookmarked: true})
	}
	return entries, nil
}

func TestCardModelMarkReadError(t *testing.T) {
	store := cardStore{markReadErr: errors.New("boom")}
	m := CardModel{
		Article: &articles.Article{Title: "t", URL: "u"},
		History: store,
	}

	model, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("m")})
	msg := cmd()
	model2, _ := model.(CardModel).Update(msg)
	cm := model2.(CardModel)

	if cm.Message != "Failed to mark as read" {
		t.Fatalf("expected error message, got %q", cm.Message)
	}
}

func TestCardModelBookmarkError(t *testing.T) {
	store := cardStore{toggleBookmarkErr: errors.New("boom")}
	m := CardModel{
		Article: &articles.Article{Title: "t", URL: "u"},
		History: store,
	}

	model, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("b")})
	msg := cmd()
	model2, _ := model.(CardModel).Update(msg)
	cm := model2.(CardModel)

	if cm.Message != "Failed to toggle bookmark" {
		t.Fatalf("expected error message, got %q", cm.Message)
	}
}

func TestCardModelStatsUnavailable(t *testing.T) {
	store := cardStore{readErr: errors.New("boom")}
	m := CardModel{History: store}

	msg := m.loadStatsCmd()()
	model2, _ := m.Update(msg)
	cm := model2.(CardModel)

	if cm.Stats != "Stats unavailable" {
		t.Fatalf("expected stats unavailable, got %q", cm.Stats)
	}
}

func TestCardModelStatsSuccess(t *testing.T) {
	store := cardStore{
		readUrls: map[string]bool{"a": true, "b": true},
		streak:   3,
	}
	m := CardModel{History: store}

	msg := m.loadStatsCmd()()
	model2, _ := m.Update(msg)
	stats := model2.(CardModel).Stats

	if stats == "" || stats == "Stats unavailable" {
		t.Fatalf("expected stats text, got %q", stats)
	}
}

func TestCardModelToggleReadUpdatesIcons(t *testing.T) {
	store := &mutableStore{
		read:      map[string]bool{},
		bookmarks: map[string]bool{},
	}
	m := CardModel{
		Article: &articles.Article{Title: "Title", URL: "http://example.com"},
		History: store,
	}

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("m")})
	cm := updated.(CardModel)
	if !cm.IsRead {
		t.Fatalf("expected IsRead true")
	}
	view := cm.View()
	if !strings.Contains(view, "✓") {
		t.Fatalf("expected read icon in view, got %q", view)
	}

	updated, _ = cm.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("m")})
	cm = updated.(CardModel)
	if cm.IsRead {
		t.Fatalf("expected IsRead false")
	}
	view = cm.View()
	if !strings.Contains(view, "○") {
		t.Fatalf("expected unread icon in view, got %q", view)
	}
}

func TestCardModelToggleBookmarkUpdatesIcons(t *testing.T) {
	store := &mutableStore{
		read:      map[string]bool{},
		bookmarks: map[string]bool{},
	}
	m := CardModel{
		Article: &articles.Article{Title: "Title", URL: "http://example.com"},
		History: store,
	}

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("b")})
	cm := updated.(CardModel)
	if !cm.IsBookmarked {
		t.Fatalf("expected IsBookmarked true")
	}
	view := cm.View()
	if !strings.Contains(view, "★") {
		t.Fatalf("expected bookmark icon in view, got %q", view)
	}

	updated, _ = cm.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("b")})
	cm = updated.(CardModel)
	if cm.IsBookmarked {
		t.Fatalf("expected IsBookmarked false")
	}
	view = cm.View()
	if strings.Contains(view, "★") {
		t.Fatalf("did not expect bookmark icon in view, got %q", view)
	}
}
