package articles

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetDataFailsAfterRetries(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gorandom-sources-fail")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	os.Setenv("GORANDOM_CONFIG_DIR", tmpDir)
	defer os.Unsetenv("GORANDOM_CONFIG_DIR")

	var calls int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		http.Error(w, "temporary failure", http.StatusInternalServerError)
	}))
	defer server.Close()

	os.Setenv("GORANDOM_SOURCES_URL", server.URL)
	defer os.Unsetenv("GORANDOM_SOURCES_URL")

	ResetCache()

	_, err = GetData()
	if err == nil {
		t.Fatal("expected error after retries")
	}
	if calls != fetchRetryCount {
		t.Fatalf("expected %d calls, got %d", fetchRetryCount, calls)
	}
}
