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

	ResetCache()

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
