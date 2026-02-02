package articles

import (
	"testing"
)

func TestDataIntegrity(t *testing.T) {
	if len(Data) == 0 {
		t.Fatal("No articles found in Data")
	}

	for i, a := range Data {
		if a.URL == "" {
			t.Errorf("Article at index %d has empty URL", i)
		}
		if a.Title == "" {
			t.Errorf("Article at index %d has empty Title", i)
		}
		if a.Source == "" {
			t.Errorf("Article at index %d has empty Source", i)
		}
	}
}

func TestFormatSource(t *testing.T) {
	tests := []struct {
		input    Source
		expected string
	}{
		{SourceTour, "Tour of Go"},
		{SourceBlog, "Go Blog"},
		{"unknown", "unknown"},
	}

	for _, tt := range tests {
		got := FormatSource(tt.input)
		if got == "" {
			t.Errorf("FormatSource(%q) returned empty string", tt.input)
		}
	}
}

func TestPickRandom(t *testing.T) {
	// Test with data
	pool := []Article{
		{Title: "A", URL: "a"},
		{Title: "B", URL: "b"},
	}

	for i := 0; i < 10; i++ {
		picked := PickRandom(pool)
		if picked == nil {
			t.Fatal("PickRandom returned nil for non-empty pool")
		}
		if picked.Title != "A" && picked.Title != "B" {
			t.Errorf("PickRandom returned unknown article: %v", picked)
		}
	}

	// Test with empty
	if PickRandom(nil) != nil {
		t.Error("PickRandom(nil) should return nil")
	}
	if PickRandom([]Article{}) != nil {
		t.Error("PickRandom([]Article{}) should return nil")
	}
}

func TestFilterBySource(t *testing.T) {
	// Setup test data
	testData := []Article{
		{Title: "T1", Source: SourceTour, URL: "u1"},
		{Title: "T2", Source: SourceTour, URL: "u2"},
		{Title: "B1", Source: SourceBlog, URL: "u3"},
	}

	tourArticles := FilterBySource(testData, SourceTour)
	if len(tourArticles) != 2 {
		t.Errorf("Expected 2 tour articles, got %d", len(tourArticles))
	}
	for _, a := range tourArticles {
		if a.Source != SourceTour {
			t.Errorf("Expected source 'tour', got '%s'", a.Source)
		}
	}

	blogArticles := FilterBySource(testData, SourceBlog)
	if len(blogArticles) != 1 {
		t.Errorf("Expected 1 blog article, got %d", len(blogArticles))
	}

	noneArticles := FilterBySource(testData, "other")
	if len(noneArticles) != 0 {
		t.Errorf("Expected 0 articles, got %d", len(noneArticles))
	}
}

func TestGetUnreadArticles(t *testing.T) {
	testData := []Article{
		{Title: "A1", URL: "url1"},
		{Title: "A2", URL: "url2"},
		{Title: "A3", URL: "url3"},
	}

	// Case 1: No read articles
	unread1 := GetUnreadArticles(testData, nil)
	if len(unread1) != 3 {
		t.Errorf("Expected 3 unread articles, got %d", len(unread1))
	}

	// Case 2: Some read articles
	readUrls := map[string]bool{
		"url1": true,
		"url3": true,
	}
	unread2 := GetUnreadArticles(testData, readUrls)
	if len(unread2) != 1 {
		t.Errorf("Expected 1 unread article, got %d", len(unread2))
	}
	if unread2[0].URL != "url2" {
		t.Errorf("Expected unread article to be url2, got %s", unread2[0].URL)
	}

	// Case 3: All read
	readUrlsAll := map[string]bool{
		"url1": true,
		"url2": true,
		"url3": true,
	}
	unread3 := GetUnreadArticles(testData, readUrlsAll)
	if len(unread3) != 0 {
		t.Errorf("Expected 0 unread articles, got %d", len(unread3))
	}
}
