package history

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

var renameFile = os.Rename

type HistoryEntry struct {
	URL          string    `json:"url"`
	ViewedAt     time.Time `json:"viewedAt"`
	IsRead       bool      `json:"isRead"`
	IsBookmarked bool      `json:"isBookmarked"`
	ReadAt       time.Time `json:"readAt"`
}

type HistoryData struct {
	Entries []HistoryEntry `json:"entries"`
}

func GetConfigDir() (string, error) {
	if env := os.Getenv("GORANDOM_CONFIG_DIR"); env != "" {
		return env, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".random-go"), nil
}

func GetHistoryFilePath() (string, error) {
	dir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "history.json"), nil
}

func LoadHistory() (*HistoryData, error) {
	path, err := GetHistoryFilePath()
	if err != nil {
		return nil, fmt.Errorf("failed to get history path: %w", err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return &HistoryData{Entries: []HistoryEntry{}}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read history file: %w", err)
	}

	var history HistoryData
	if err := json.Unmarshal(data, &history); err != nil {
		return nil, fmt.Errorf("failed to parse history file: %w", err)
	}

	return &history, nil
}

func SaveHistory(history *HistoryData) error {
	dir, err := GetConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config dir: %w", err)
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config dir: %w", err)
	}

	path, err := GetHistoryFilePath()
	if err != nil {
		return fmt.Errorf("failed to get history path: %w", err)
	}

	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal history: %w", err)
	}

	tmp, err := os.CreateTemp(dir, "history-*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temp history file: %w", err)
	}
	tmpPath := tmp.Name()
	defer func() {
		if tmpPath != "" {
			_ = os.Remove(tmpPath)
		}
	}()

	if _, err := tmp.Write(data); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("failed to write temp history file: %w", err)
	}
	if err := tmp.Sync(); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("failed to sync temp history file: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("failed to close temp history file: %w", err)
	}

	if err := renameFile(tmpPath, path); err != nil {
		return fmt.Errorf("failed to move temp history file into place: %w", err)
	}
	tmpPath = ""
	return nil
}

// FindByURL returns the index and pointer to the entry with the given URL.
// Returns -1 and nil if not found.
func (h *HistoryData) FindByURL(url string) (int, *HistoryEntry) {
	for i := range h.Entries {
		if h.Entries[i].URL == url {
			return i, &h.Entries[i]
		}
	}
	return -1, nil
}

func (s *FileStore) AddToHistory(url string) error {
	history, err := LoadHistory()
	if err != nil {
		return err
	}

	_, entry := history.FindByURL(url)
	if entry != nil {
		entry.ViewedAt = time.Now()
	} else {
		history.Entries = append(history.Entries, HistoryEntry{
			URL:      url,
			ViewedAt: time.Now(),
			IsRead:   false,
		})
	}

	return SaveHistory(history)
}

func (s *FileStore) MarkAsRead(url string) error {
	history, err := LoadHistory()
	if err != nil {
		return err
	}

	_, entry := history.FindByURL(url)
	if entry != nil {
		// Set ReadAt if this is the first time marking as read
		if !entry.IsRead {
			entry.ReadAt = time.Now()
		}
		entry.IsRead = true
	} else {
		history.Entries = append(history.Entries, HistoryEntry{
			URL:      url,
			ViewedAt: time.Now(),
			IsRead:   true,
			ReadAt:   time.Now(),
		})
	}

	return SaveHistory(history)
}

func (s *FileStore) MarkAsUnread(url string) (bool, error) {
	history, err := LoadHistory()
	if err != nil {
		return false, err
	}

	_, entry := history.FindByURL(url)
	if entry != nil {
		entry.IsRead = false
		err := SaveHistory(history)
		return true, err
	}
	return false, nil
}

func (s *FileStore) GetReadUrls() (map[string]bool, error) {
	history, err := LoadHistory()
	if err != nil {
		return nil, err
	}

	readUrls := make(map[string]bool)
	for _, entry := range history.Entries {
		if entry.IsRead {
			readUrls[entry.URL] = true
		}
	}
	return readUrls, nil
}

func (s *FileStore) GetSortedHistory() ([]HistoryEntry, error) {
	history, err := LoadHistory()
	if err != nil {
		return nil, err
	}

	// Sort by ViewedAt descending (newest first)
	entries := make([]HistoryEntry, len(history.Entries))
	copy(entries, history.Entries)

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].ViewedAt.After(entries[j].ViewedAt)
	})

	return entries, nil
}

// Function wrappers for backward compatibility
func AddToHistory(url string) error             { return DefaultStore.AddToHistory(url) }
func MarkAsRead(url string) error               { return DefaultStore.MarkAsRead(url) }
func MarkAsUnread(url string) (bool, error)     { return DefaultStore.MarkAsUnread(url) }
func GetReadUrls() (map[string]bool, error)     { return DefaultStore.GetReadUrls() }
func GetSortedHistory() ([]HistoryEntry, error) { return DefaultStore.GetSortedHistory() }
