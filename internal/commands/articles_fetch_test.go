package commands

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/AfshinJalili/gorandom/internal/articles"
	"github.com/spf13/cobra"
)

func TestFetchMessageShownOnce(t *testing.T) {
	resetSourcesFetchOnce()

	disableSpinnerForTests(t)

	tmpDir, err := os.MkdirTemp("", "gorandom-fetch-msg")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	os.Setenv("GORANDOM_CONFIG_DIR", tmpDir)
	defer os.Unsetenv("GORANDOM_CONFIG_DIR")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := struct {
			Version   int                `json:"version"`
			UpdatedAt time.Time          `json:"updatedAt"`
			Articles  []articles.Article `json:"articles"`
		}{
			Version:   1,
			UpdatedAt: time.Now().UTC(),
			Articles: []articles.Article{
				{URL: "http://example.com", Title: "Example", Source: articles.SourceDocs},
			},
		}
		_ = json.NewEncoder(w).Encode(payload)
	}))
	defer server.Close()

	os.Setenv("GORANDOM_SOURCES_URL", server.URL)
	defer os.Unsetenv("GORANDOM_SOURCES_URL")

	articles.Data = nil

	cmd := &cobra.Command{}
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)

	_, ok := loadArticlesRequired(cmd)
	if !ok {
		t.Fatal("expected loadArticlesRequired success")
	}
	if !bytes.Contains(buf.Bytes(), []byte("Fetching sources...")) {
		t.Fatalf("expected fetch message, got %q", buf.String())
	}

	buf.Reset()
	cmd.SetOut(buf)
	_, ok = loadArticlesRequired(cmd)
	if !ok {
		t.Fatal("expected loadArticlesRequired success on second call")
	}
	if bytes.Contains(buf.Bytes(), []byte("Fetching sources...")) {
		t.Fatalf("did not expect fetch message on second call, got %q", buf.String())
	}
}

func disableSpinnerForTests(t *testing.T) {
	t.Helper()
	os.Setenv("GORANDOM_SOURCES_SPINNER", "0")
	t.Cleanup(func() {
		os.Unsetenv("GORANDOM_SOURCES_SPINNER")
	})
}
