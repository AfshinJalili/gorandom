package output

import "time"

type Article struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	Source string `json:"source"`
}

type HistoryEntry struct {
	Title        string    `json:"title"`
	URL          string    `json:"url"`
	ViewedAt     time.Time `json:"viewedAt"`
	IsRead       bool      `json:"isRead"`
	IsBookmarked bool      `json:"isBookmarked"`
}
