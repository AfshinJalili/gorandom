package history

// Store defines the interface for history management.
type Store interface {
	AddToHistory(url string) error
	MarkAsRead(url string) error
	MarkAsUnread(url string) (bool, error)
	GetReadUrls() (map[string]bool, error)
	GetSortedHistory() ([]HistoryEntry, error)
	CalculateStreak() (int, error)
	ToggleBookmark(url string) (bool, error)
	GetBookmarks() ([]HistoryEntry, error)
}

// FileStore is an implementation of Store that uses the filesystem.
type FileStore struct{}

func NewFileStore() *FileStore {
	return &FileStore{}
}

// DefaultStore is a global instance for convenience and backward compatibility.
var DefaultStore Store = NewFileStore()
