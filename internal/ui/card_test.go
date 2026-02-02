package ui

import (
	"errors"
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
