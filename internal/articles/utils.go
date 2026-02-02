package articles

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func PickRandom(pool []Article) *Article {
	if len(pool) == 0 {
		return nil
	}
	return &pool[rand.Intn(len(pool))]
}

func FilterBySource(pool []Article, source Source) []Article {
	var filtered []Article
	for _, a := range pool {
		if a.Source == source {
			filtered = append(filtered, a)
		}
	}
	return filtered
}

func GetUnreadArticles(pool []Article, readUrls map[string]bool) []Article {
	var unread []Article
	for _, a := range pool {
		if !readUrls[a.URL] {
			unread = append(unread, a)
		}
	}
	return unread
}
