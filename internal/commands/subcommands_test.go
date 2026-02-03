package commands

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/AfshinJalili/gorandom/internal/articles"
	"github.com/AfshinJalili/gorandom/internal/history"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func setupTestEnv(t *testing.T) string {
	tmpDir, err := os.MkdirTemp("", "gorandom-cmd-test")
	if err != nil {
		t.Fatal(err)
	}
	os.Setenv("GORANDOM_CONFIG_DIR", tmpDir)
	os.Setenv("GORANDOM_SOURCES_AUTO_UPDATE", "0")
	if len(articles.Data) == 0 {
		articles.Data = []articles.Article{
			{URL: "http://example.com/default", Source: articles.SourceDocs, Title: "Default"},
		}
	}
	return tmpDir
}

func resetFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		_ = cmd.Flags().Set(f.Name, f.DefValue)
	})
	cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		_ = cmd.PersistentFlags().Set(f.Name, f.DefValue)
	})
	for _, sub := range cmd.Commands() {
		resetFlags(sub)
	}
}

func TestSourcesCommand(t *testing.T) {
	setupTestEnv(t)
	// No need to cleanup env var as it's global, but tmpDir should be removed
	// Actually better to have a cleanup helper
}

func TestCommandSuite(t *testing.T) {
	tmpDir := setupTestEnv(t)
	defer os.RemoveAll(tmpDir)
	defer os.Unsetenv("GORANDOM_CONFIG_DIR")
	defer os.Unsetenv("GORANDOM_SOURCES_AUTO_UPDATE")
	resetSourcesFetchOnce()

	t.Run("Sources", func(t *testing.T) {
		resetFlags(rootCmd)
		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"sources"})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), "Available sources:") {
			t.Errorf("Expected 'Available sources:', got %q", b.String())
		}
	})

	t.Run("BookmarkURL", func(t *testing.T) {
		resetFlags(rootCmd)
		url := "http://example.com/test-bookmark"
		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"bookmark", url})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), "Bookmarked") {
			t.Errorf("Expected 'Bookmarked', got %q", b.String())
		}

		// Toggle off
		b.Reset()
		rootCmd.SetArgs([]string{"bookmark", url})
		rootCmd.Execute()
		if !strings.Contains(b.String(), "Removed bookmark") {
			t.Errorf("Expected 'Removed bookmark', got %q", b.String())
		}
	})

	t.Run("MarkURL", func(t *testing.T) {
		resetFlags(rootCmd)
		url := "http://example.com/test-mark"
		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"mark", url})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), "Marked as read") {
			t.Errorf("Expected 'Marked as read', got %q", b.String())
		}

		// Verify in history
		read, _ := history.GetReadUrls()
		if !read[url] {
			t.Error("Expected URL to be read in history")
		}

		// Unmark
		b.Reset()
		rootCmd.SetArgs([]string{"unmark", url})
		rootCmd.Execute()
		if !strings.Contains(b.String(), "Marked as unread") {
			t.Errorf("Expected 'Marked as unread', got %q", b.String())
		}
	})

	t.Run("MarkIndex", func(t *testing.T) {
		resetFlags(rootCmd)
		url := "http://example.com/index-test"
		history.AddToHistory(url) // index 1

		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"mark", "1"})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), "Marked as read") {
			t.Errorf("Expected 'Marked as read', got %q", b.String())
		}
	})

	t.Run("InvalidSource", func(t *testing.T) {
		resetFlags(rootCmd)
		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"random", "--source", "invalid"})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), "Invalid source") {
			t.Errorf("Expected 'Invalid source', got %q", b.String())
		}
		if !strings.Contains(b.String(), "Try:") {
			t.Errorf("Expected hint line, got %q", b.String())
		}
	})

	t.Run("OutputModeConflict", func(t *testing.T) {
		resetFlags(rootCmd)
		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"random", "--plain", "--json"})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), "use only one") {
			t.Errorf("Expected output mode error, got %q", b.String())
		}
		if !strings.Contains(b.String(), "Try:") {
			t.Errorf("Expected hint line, got %q", b.String())
		}
	})

	t.Run("RandomPlain", func(t *testing.T) {
		resetFlags(rootCmd)
		orig := articles.Data
		articles.Data = []articles.Article{
			{URL: "http://example.com/plain", Source: articles.SourceDocs, Title: "Plain Title"},
		}
		defer func() { articles.Data = orig }()

		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"random", "--plain", "--any"})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), "Plain Title") || !strings.Contains(b.String(), "http://example.com/plain") {
			t.Errorf("Unexpected output: %q", b.String())
		}
	})

	t.Run("RandomJSON", func(t *testing.T) {
		resetFlags(rootCmd)
		orig := articles.Data
		articles.Data = []articles.Article{
			{URL: "http://example.com/json", Source: articles.SourceDocs, Title: "JSON Title"},
		}
		defer func() { articles.Data = orig }()

		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"random", "--json", "--any"})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), "\"url\"") || !strings.Contains(b.String(), "http://example.com/json") {
			t.Errorf("Unexpected JSON output: %q", b.String())
		}
	})

	t.Run("HistoryPlain", func(t *testing.T) {
		resetFlags(rootCmd)
		url := "http://example.com/history-plain"
		if err := history.AddToHistory(url); err != nil {
			t.Fatal(err)
		}

		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"history", "--plain", "--limit", "1"})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), url) {
			t.Errorf("Unexpected output: %q", b.String())
		}
	})

	t.Run("HistoryJSON", func(t *testing.T) {
		resetFlags(rootCmd)
		url := "http://example.com/history-json"
		if err := history.AddToHistory(url); err != nil {
			t.Fatal(err)
		}

		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"history", "--json", "--limit", "1"})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), "\"url\"") || !strings.Contains(b.String(), url) {
			t.Errorf("Unexpected JSON output: %q", b.String())
		}
	})

	t.Run("NextPlain", func(t *testing.T) {
		resetFlags(rootCmd)
		orig := articles.Data
		articles.Data = []articles.Article{
			{URL: "http://example.com/next", Source: articles.SourceDocs, Title: "Next Title"},
		}
		defer func() { articles.Data = orig }()

		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"next", "--plain"})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), "Next Title") || !strings.Contains(b.String(), "http://example.com/next") {
			t.Errorf("Unexpected output: %q", b.String())
		}
	})

	t.Run("NextJSON", func(t *testing.T) {
		resetFlags(rootCmd)
		orig := articles.Data
		articles.Data = []articles.Article{
			{URL: "http://example.com/next-json", Source: articles.SourceDocs, Title: "Next JSON"},
		}
		defer func() { articles.Data = orig }()

		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"next", "--json"})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), "\"url\"") || !strings.Contains(b.String(), "http://example.com/next-json") {
			t.Errorf("Unexpected JSON output: %q", b.String())
		}
	})

	t.Run("SearchPlain", func(t *testing.T) {
		resetFlags(rootCmd)
		orig := articles.Data
		articles.Data = []articles.Article{
			{URL: "http://example.com/alpha", Source: articles.SourceDocs, Title: "Alpha Guide"},
			{URL: "http://example.com/beta", Source: articles.SourceBlog, Title: "Beta Tips"},
		}
		defer func() { articles.Data = orig }()

		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"search", "Alpha", "--plain"})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), "Alpha Guide") || !strings.Contains(b.String(), "http://example.com/alpha") {
			t.Errorf("Unexpected output: %q", b.String())
		}
	})

	t.Run("SearchJSON", func(t *testing.T) {
		resetFlags(rootCmd)
		orig := articles.Data
		articles.Data = []articles.Article{
			{URL: "http://example.com/alpha-json", Source: articles.SourceDocs, Title: "Alpha JSON"},
		}
		defer func() { articles.Data = orig }()

		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"search", "Alpha", "--json"})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), "\"url\"") || !strings.Contains(b.String(), "http://example.com/alpha-json") {
			t.Errorf("Unexpected JSON output: %q", b.String())
		}
	})

	t.Run("HistoryCorrupted", func(t *testing.T) {
		resetFlags(rootCmd)
		path, err := history.GetHistoryFilePath()
		if err != nil {
			t.Fatal(err)
		}
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte("{bad json"), 0644); err != nil {
			t.Fatal(err)
		}

		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"history", "--plain"})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), "Could not load history") {
			t.Errorf("Expected history load error, got %q", b.String())
		}
		if !strings.Contains(b.String(), "Try:") {
			t.Errorf("Expected hint line, got %q", b.String())
		}
	})

	t.Run("OpenInvalidIndex", func(t *testing.T) {
		resetFlags(rootCmd)
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Fatal(err)
		}
		if err := os.MkdirAll(tmpDir, 0755); err != nil {
			t.Fatal(err)
		}
		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"open", "9999"})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), "Invalid index") {
			t.Errorf("Expected invalid index error, got %q", b.String())
		}
		if !strings.Contains(b.String(), "Try:") {
			t.Errorf("Expected hint line, got %q", b.String())
		}
	})

	t.Run("OpenHistoryError", func(t *testing.T) {
		resetFlags(rootCmd)
		path, err := history.GetHistoryFilePath()
		if err != nil {
			t.Fatal(err)
		}
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte("{bad json"), 0644); err != nil {
			t.Fatal(err)
		}

		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"open", "1"})
		if err := rootCmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(b.String(), "Could not load history") {
			t.Errorf("Expected history load error, got %q", b.String())
		}
		if !strings.Contains(b.String(), "Try:") {
			t.Errorf("Expected hint line, got %q", b.String())
		}
	})
}
