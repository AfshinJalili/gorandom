package articles

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestGetDataRetriesOnMissingCache(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gorandom-sources-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	os.Setenv("GORANDOM_CONFIG_DIR", tmpDir)
	defer os.Unsetenv("GORANDOM_CONFIG_DIR")

	var calls int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if calls == 1 {
			http.Error(w, "temporary failure", http.StatusInternalServerError)
			return
		}
		payload := sourcesFile{
			Version:   1,
			UpdatedAt: time.Now().UTC(),
			Articles: []Article{
				{URL: "http://example.com", Title: "Example", Source: SourceDocs},
			},
		}
		_ = json.NewEncoder(w).Encode(payload)
	}))
	defer server.Close()

	os.Setenv("GORANDOM_SOURCES_URL", server.URL)
	defer os.Unsetenv("GORANDOM_SOURCES_URL")

	dataMu.Lock()
	Data = nil
	dataLoaded = false
	dataMu.Unlock()

	data, err := GetData()
	if err != nil {
		t.Fatalf("expected GetData success, got %v", err)
	}
	if len(data) != 1 {
		t.Fatalf("expected 1 article, got %d", len(data))
	}
	if calls != fetchRetryCount {
		t.Fatalf("expected %d calls, got %d", fetchRetryCount, calls)
	}
}

func TestWriteCacheAtomicRenameFailureKeepsOriginalFile(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gorandom-cache-atomic")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	os.Setenv("GORANDOM_CONFIG_DIR", tmpDir)
	defer os.Unsetenv("GORANDOM_CONFIG_DIR")

	cachePath, err := getCachePath()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		t.Fatal(err)
	}

	original := []byte("{\n  \"version\": 1,\n  \"updatedAt\": \"2024-01-01T00:00:00Z\",\n  \"articles\": [{\"url\":\"http://existing\",\"title\":\"Existing\",\"source\":\"docs\"}]\n}")
	if err := os.WriteFile(cachePath, original, 0644); err != nil {
		t.Fatal(err)
	}

	oldRename := renameCacheFile
	renameCacheFile = func(oldPath, newPath string) error { return os.ErrPermission }
	defer func() { renameCacheFile = oldRename }()

	err = writeCache([]Article{{URL: "http://new", Title: "New", Source: SourceDocs}})
	if err == nil {
		t.Fatal("expected writeCache to fail")
	}

	after, err := os.ReadFile(cachePath)
	if err != nil {
		t.Fatal(err)
	}
	if string(after) != string(original) {
		t.Fatal("expected original cache file to remain unchanged")
	}

	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		if len(file.Name()) >= 7 && file.Name()[:7] == "atomic-" {
			t.Fatalf("expected temp file cleanup, found %s", file.Name())
		}
	}
}

func TestWriteMetaAtomicRenameFailureKeepsOriginalFile(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gorandom-meta-atomic")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	os.Setenv("GORANDOM_CONFIG_DIR", tmpDir)
	defer os.Unsetenv("GORANDOM_CONFIG_DIR")

	metaPath, err := getMetaPath()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		t.Fatal(err)
	}

	original := []byte("{\n  \"fetchedAt\": \"2024-01-01T00:00:00Z\",\n  \"etag\": \"abc\",\n  \"url\": \"http://example.com\"\n}")
	if err := os.WriteFile(metaPath, original, 0644); err != nil {
		t.Fatal(err)
	}

	oldRename := renameCacheFile
	renameCacheFile = func(oldPath, newPath string) error { return os.ErrPermission }
	defer func() { renameCacheFile = oldRename }()

	err = writeMeta(cacheMeta{FetchedAt: time.Now(), Etag: "new", URL: "http://new"})
	if err == nil {
		t.Fatal("expected writeMeta to fail")
	}

	after, err := os.ReadFile(metaPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(after) != string(original) {
		t.Fatal("expected original meta file to remain unchanged")
	}

	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		if len(file.Name()) >= 7 && file.Name()[:7] == "atomic-" {
			t.Fatalf("expected temp file cleanup, found %s", file.Name())
		}
	}
}
