package commands

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/AfshinJalili/gorandom/internal/articles"
	"github.com/spf13/cobra"
)

func TestLoadArticlesRequiredReportsFailure(t *testing.T) {
	resetSourcesFetchOnce()
	disableSpinnerForTests(t)

	tmpDir, err := os.MkdirTemp("", "gorandom-fetch-fail")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	os.Setenv("GORANDOM_CONFIG_DIR", tmpDir)
	defer os.Unsetenv("GORANDOM_CONFIG_DIR")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "temporary failure", http.StatusInternalServerError)
	}))
	defer server.Close()

	os.Setenv("GORANDOM_SOURCES_URL", server.URL)
	defer os.Unsetenv("GORANDOM_SOURCES_URL")

	articles.ResetCache()

	cmd := &cobra.Command{}
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)

	_, ok := loadArticlesRequired(cmd)
	if ok {
		t.Fatal("expected loadArticlesRequired to fail")
	}

	output := buf.String()
	if !bytes.Contains(buf.Bytes(), []byte("Could not load sources:")) {
		t.Fatalf("expected error output, got %q", output)
	}
	if !bytes.Contains(buf.Bytes(), []byte("Try:")) {
		t.Fatalf("expected hint output, got %q", output)
	}
}
