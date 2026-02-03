package articles

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/AfshinJalili/gorandom/internal/history"
)

const (
	defaultSourcesURL = "https://raw.githubusercontent.com/AfshinJalili/gorandom/main/data/sources.json"
	defaultTTL        = 24 * time.Hour
)

var (
	cacheOnce sync.Once
	cacheErr  error
)

type sourcesFile struct {
	Version   int       `json:"version"`
	UpdatedAt time.Time `json:"updatedAt"`
	Articles  []Article `json:"articles"`
}

type cacheMeta struct {
	FetchedAt time.Time `json:"fetchedAt"`
	Etag      string    `json:"etag"`
	URL       string    `json:"url"`
}

type CacheStatus struct {
	CachePath string
	MetaPath  string
	URL       string
	FetchedAt time.Time
	Age       time.Duration
	Stale     bool
}

func GetSourcesURL() string {
	if v := os.Getenv("GORANDOM_SOURCES_URL"); v != "" {
		return v
	}
	return defaultSourcesURL
}

func getTTL() time.Duration {
	if v := os.Getenv("GORANDOM_SOURCES_TTL"); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return defaultTTL
}

func autoUpdateEnabled() bool {
	if v := os.Getenv("GORANDOM_SOURCES_AUTO_UPDATE"); v != "" {
		enabled, err := strconv.ParseBool(v)
		if err == nil {
			return enabled
		}
		if v == "0" {
			return false
		}
	}
	return true
}

func GetData() ([]Article, error) {
	if err := ensureCacheLoaded(); err != nil {
		return Data, err
	}
	return Data, nil
}

func ensureCacheLoaded() error {
	cacheOnce.Do(func() {
		cacheErr = loadCache()
	})
	return cacheErr
}

func loadCache() error {
	cachePath, err := getCachePath()
	if err != nil {
		return err
	}
	data, err := os.ReadFile(cachePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("failed to read sources cache: %w", err)
	}

	var parsed sourcesFile
	if err := json.Unmarshal(data, &parsed); err != nil {
		return fmt.Errorf("failed to parse sources cache: %w", err)
	}
	if len(parsed.Articles) == 0 {
		return fmt.Errorf("sources cache contains no articles")
	}
	Data = parsed.Articles
	return nil
}

func SyncIfStale() (bool, error) {
	if !autoUpdateEnabled() {
		return false, nil
	}
	ttl := getTTL()
	stale, err := isCacheStale(ttl)
	if err != nil {
		return false, err
	}
	if !stale {
		return false, nil
	}
	return updateFromRemote(false)
}

func ForceUpdate() (bool, error) {
	return updateFromRemote(true)
}

func CacheStatusInfo() (CacheStatus, error) {
	cachePath, err := getCachePath()
	if err != nil {
		return CacheStatus{}, err
	}
	metaPath, err := getMetaPath()
	if err != nil {
		return CacheStatus{}, err
	}
	meta, _ := readMeta()
	age := time.Since(meta.FetchedAt)
	stale := meta.FetchedAt.IsZero() || age > getTTL()
	return CacheStatus{
		CachePath: cachePath,
		MetaPath:  metaPath,
		URL:       GetSourcesURL(),
		FetchedAt: meta.FetchedAt,
		Age:       age,
		Stale:     stale,
	}, nil
}

func updateFromRemote(force bool) (bool, error) {
	if err := ensureCacheLoaded(); err != nil {
		// Continue with remote fetch; cache may be corrupt.
	}
	meta, _ := readMeta()
	if !force && !meta.FetchedAt.IsZero() && time.Since(meta.FetchedAt) < getTTL() {
		return false, nil
	}

	updated, articles, etag, err := fetchRemote(meta.Etag)
	if err != nil {
		return false, err
	}
	if !updated {
		meta.FetchedAt = time.Now()
		_ = writeMeta(meta)
		return false, nil
	}

	Data = articles
	newMeta := cacheMeta{
		FetchedAt: time.Now(),
		Etag:      etag,
		URL:       GetSourcesURL(),
	}
	if err := writeCache(articles); err != nil {
		return true, err
	}
	if err := writeMeta(newMeta); err != nil {
		return true, err
	}
	return true, nil
}

func fetchRemote(etag string) (bool, []Article, string, error) {
	req, err := http.NewRequest(http.MethodGet, GetSourcesURL(), nil)
	if err != nil {
		return false, nil, "", err
	}
	if etag != "" {
		req.Header.Set("If-None-Match", etag)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, nil, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		return false, nil, etag, nil
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return false, nil, "", fmt.Errorf("sources fetch failed: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, nil, "", err
	}
	var parsed sourcesFile
	if err := json.Unmarshal(body, &parsed); err != nil {
		return false, nil, "", fmt.Errorf("invalid sources payload: %w", err)
	}
	if len(parsed.Articles) == 0 {
		return false, nil, "", fmt.Errorf("sources payload contains no articles")
	}
	return true, parsed.Articles, resp.Header.Get("ETag"), nil
}

func isCacheStale(ttl time.Duration) (bool, error) {
	meta, err := readMeta()
	if err != nil {
		return true, nil
	}
	if meta.FetchedAt.IsZero() {
		return true, nil
	}
	return time.Since(meta.FetchedAt) > ttl, nil
}

func getCachePath() (string, error) {
	dir, err := history.GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "sources.json"), nil
}

func getMetaPath() (string, error) {
	dir, err := history.GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "sources.meta.json"), nil
}

func writeCache(articles []Article) error {
	cachePath, err := getCachePath()
	if err != nil {
		return err
	}
	dir := filepath.Dir(cachePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	payload := sourcesFile{
		Version:   1,
		UpdatedAt: time.Now().UTC(),
		Articles:  articles,
	}
	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(cachePath, data, 0644)
}

func readMeta() (cacheMeta, error) {
	metaPath, err := getMetaPath()
	if err != nil {
		return cacheMeta{}, err
	}
	data, err := os.ReadFile(metaPath)
	if err != nil {
		return cacheMeta{}, err
	}
	var meta cacheMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return cacheMeta{}, err
	}
	return meta, nil
}

func writeMeta(meta cacheMeta) error {
	metaPath, err := getMetaPath()
	if err != nil {
		return err
	}
	dir := filepath.Dir(metaPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(metaPath, data, 0644)
}
