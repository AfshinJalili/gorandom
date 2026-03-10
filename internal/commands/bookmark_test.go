package commands

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/AfshinJalili/gorandom/internal/history"
	"github.com/spf13/cobra"
)

type failingHistoryStore struct {
	history.Store
	getSortedHistoryErr error
}

func (s failingHistoryStore) AddToHistory(url string) error { return nil }

func (s failingHistoryStore) MarkAsRead(url string) error { return nil }

func (s failingHistoryStore) MarkAsUnread(url string) (bool, error) { return false, nil }

func (s failingHistoryStore) GetReadUrls() (map[string]bool, error) { return map[string]bool{}, nil }

func (s failingHistoryStore) GetSortedHistory() ([]history.HistoryEntry, error) {
	return nil, s.getSortedHistoryErr
}

func (s failingHistoryStore) CalculateStreak() (int, error) { return 0, nil }

func (s failingHistoryStore) ToggleBookmark(url string) (bool, error) { return false, nil }

func (s failingHistoryStore) GetBookmarks() ([]history.HistoryEntry, error) { return nil, nil }

func TestRunBookmark_IndexHistoryLoadError(t *testing.T) {
	origStore := historyStore
	historyStore = failingHistoryStore{getSortedHistoryErr: errors.New("boom")}
	t.Cleanup(func() {
		historyStore = origStore
	})

	cmd := &cobra.Command{}
	buf := bytes.NewBuffer(nil)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	runBookmark(cmd, []string{"1"})

	out := buf.String()
	if !strings.Contains(out, "Could not load history: boom") {
		t.Fatalf("expected history load error, got %q", out)
	}
	if strings.Contains(out, "Invalid URL or index") {
		t.Fatalf("expected no invalid input error for storage failure, got %q", out)
	}
}
