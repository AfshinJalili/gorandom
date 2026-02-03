package articles

// Cached returns the in-memory articles slice without triggering a fetch.
func Cached() []Article {
	dataMu.RLock()
	defer dataMu.RUnlock()
	return cachedArticles
}

// SetCached replaces the in-memory articles slice (used by tests).
func SetCached(articles []Article) {
	dataMu.Lock()
	cachedArticles = articles
	dataLoaded = len(articles) > 0
	dataMu.Unlock()
}

// ResetCache clears the in-memory cache (used by tests).
func ResetCache() {
	SetCached(nil)
}
