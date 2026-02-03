package reader

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	nurl "net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/AfshinJalili/gorandom/internal/history"
	readability "github.com/go-shiori/go-readability"
)

type Content struct {
	URL        string    `json:"url"`
	Title      string    `json:"title"`
	Paragraphs []string  `json:"paragraphs"`
	FetchedAt  time.Time `json:"fetchedAt"`
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

func Load(url string, refresh bool) (*Content, error) {
	path, err := cachePath(url)
	if err != nil {
		return nil, err
	}

	if !refresh {
		if cached, err := readCache(path); err == nil {
			return cached, nil
		}
	}

	body, err := fetchHTML(url)
	if err != nil {
		return nil, err
	}

	title, paragraphs, err := extractReadable(url, body)
	if err != nil {
		return nil, err
	}

	content := &Content{
		URL:        url,
		Title:      title,
		Paragraphs: paragraphs,
		FetchedAt:  time.Now().UTC(),
	}

	if err := writeCache(path, content); err != nil {
		return content, err
	}
	return content, nil
}

func fetchHTML(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "gorandom/zen")
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("fetch failed: %s", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func cachePath(url string) (string, error) {
	dir, err := history.GetConfigDir()
	if err != nil {
		return "", err
	}
	hash := sha1.Sum([]byte(url))
	name := hex.EncodeToString(hash[:]) + ".json"
	return filepath.Join(dir, "reader", name), nil
}

func readCache(path string) (*Content, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var content Content
	if err := json.Unmarshal(data, &content); err != nil {
		return nil, err
	}
	if len(content.Paragraphs) == 0 {
		return nil, fmt.Errorf("cached content empty")
	}
	return &content, nil
}

func writeCache(path string, content *Content) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func extractTitle(htmlText string) string {
	re := regexp.MustCompile(`(?is)<title[^>]*>(.*?)</title>`)
	match := re.FindStringSubmatch(htmlText)
	if len(match) < 2 {
		return ""
	}
	return strings.TrimSpace(stripTags(match[1]))
}

func extractParagraphs(htmlText string) []string {
	re := regexp.MustCompile(`(?is)<(p|li|pre)[^>]*>(.*?)</\\1>`)
	matches := re.FindAllStringSubmatch(htmlText, -1)
	var paragraphs []string
	for _, match := range matches {
		if len(match) < 3 {
			continue
		}
		text := stripTags(match[2])
		text = normalizeWhitespace(text)
		if text != "" {
			paragraphs = append(paragraphs, text)
		}
	}

	if len(paragraphs) == 0 {
		fallback := normalizeWhitespace(stripTags(htmlText))
		if fallback != "" {
			paragraphs = append(paragraphs, fallback)
		}
	}

	return paragraphs
}

func extractReadable(url string, body string) (string, []string, error) {
	parsedURL, err := nurl.Parse(url)
	if err != nil {
		return "", nil, err
	}

	article, err := readability.FromReader(strings.NewReader(body), parsedURL)
	if err == nil {
		title := strings.TrimSpace(article.Title)
		paragraphs := extractParagraphs(article.Content)
		if len(paragraphs) == 0 {
			paragraphs = splitTextContent(article.TextContent)
		}
		if len(paragraphs) > 0 {
			if title == "" {
				title = url
			}
			return title, paragraphs, nil
		}
	}

	title := extractTitle(body)
	paragraphs := extractParagraphs(body)
	if len(paragraphs) == 0 {
		return "", nil, fmt.Errorf("no readable content found")
	}
	if title == "" {
		title = url
	}
	return title, paragraphs, nil
}

func stripTags(input string) string {
	input = strings.ReplaceAll(input, "<br>", "\n")
	input = strings.ReplaceAll(input, "<br/>", "\n")
	input = strings.ReplaceAll(input, "<br />", "\n")
	re := regexp.MustCompile(`(?is)<[^>]+>`)
	stripped := re.ReplaceAllString(input, " ")
	return html.UnescapeString(stripped)
}

func normalizeWhitespace(input string) string {
	input = strings.ReplaceAll(input, "\n", " ")
	parts := strings.Fields(input)
	return strings.TrimSpace(strings.Join(parts, " "))
}

func splitTextContent(text string) []string {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	lines := strings.Split(text, "\n")
	var paragraphs []string
	var buf []string
	flush := func() {
		if len(buf) == 0 {
			return
		}
		paragraph := normalizeWhitespace(strings.Join(buf, " "))
		if paragraph != "" {
			paragraphs = append(paragraphs, paragraph)
		}
		buf = nil
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			flush()
			continue
		}
		buf = append(buf, line)
	}
	flush()

	if len(paragraphs) == 0 {
		normalized := normalizeWhitespace(text)
		if normalized != "" {
			paragraphs = append(paragraphs, normalized)
		}
	}
	return paragraphs
}
