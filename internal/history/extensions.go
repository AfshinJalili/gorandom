package history

import (
	"sort"
	"time"
)

func (s *FileStore) ToggleBookmark(url string) (bool, error) {
	history, err := LoadHistory()
	if err != nil {
		return false, err
	}

	_, entry := history.FindByURL(url)
	var isBookmarked bool

	if entry != nil {
		entry.IsBookmarked = !entry.IsBookmarked
		isBookmarked = entry.IsBookmarked
	} else {
		// If toggling bookmark on unseen url, add it
		history.Entries = append(history.Entries, HistoryEntry{
			URL:          url,
			ViewedAt:     time.Now(),
			IsBookmarked: true,
		})
		isBookmarked = true
	}

	err = SaveHistory(history)
	return isBookmarked, err
}

func (s *FileStore) GetBookmarks() ([]HistoryEntry, error) {
	history, err := LoadHistory()
	if err != nil {
		return nil, err
	}

	var bookmarks []HistoryEntry
	for _, entry := range history.Entries {
		if entry.IsBookmarked {
			bookmarks = append(bookmarks, entry)
		}
	}
	return bookmarks, nil
}

func (s *FileStore) CalculateStreak() (int, error) {
	history, err := LoadHistory()
	if err != nil {
		return 0, err
	}

	// Filter read entries with valid ReadAt
	var readDates []time.Time
	for _, entry := range history.Entries {
		if entry.IsRead && !entry.ReadAt.IsZero() {
			readDates = append(readDates, entry.ReadAt)
		}
	}

	if len(readDates) == 0 {
		return 0, nil
	}

	// Sort dates descending
	sort.Slice(readDates, func(i, j int) bool {
		return readDates[i].After(readDates[j])
	})

	streak := 0

	// Normalize to start of day
	truncate := func(t time.Time) time.Time {
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	}

	today := truncate(time.Now())
	lastDate := truncate(readDates[0])

	diff := today.Sub(lastDate).Hours() / 24
	if diff > 1 {
		return 0, nil
	}

	currentCheck := lastDate
	processed := make(map[string]bool)
	processed[lastDate.Format("2006-01-02")] = true
	streak = 1

	for _, d := range readDates[1:] {
		date := truncate(d)
		if processed[date.Format("2006-01-02")] {
			continue
		}

		expectedPrev := currentCheck.AddDate(0, 0, -1)

		// If this date is the expected previous day
		if date.Equal(expectedPrev) {
			streak++
			currentCheck = date
			processed[date.Format("2006-01-02")] = true
		} else {
			// gap found
			break
		}
	}

	return streak, nil
}

// Function wrappers for backward compatibility
func ToggleBookmark(url string) (bool, error) { return DefaultStore.ToggleBookmark(url) }
func GetBookmarks() ([]HistoryEntry, error)   { return DefaultStore.GetBookmarks() }
func CalculateStreak() (int, error)           { return DefaultStore.CalculateStreak() }
