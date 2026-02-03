package ui

import (
	"fmt"
	"time"

	"github.com/AfshinJalili/gorandom/internal/articles"
	"github.com/AfshinJalili/gorandom/internal/history"
	tea "github.com/charmbracelet/bubbletea"
)

type CardModel struct {
	Article      *articles.Article
	Pool         []articles.Article
	Message      string
	MsgType      string // "info", "error", "success"
	Quitting     bool
	History      history.Store
	ShowHelp     bool
	Stats        string
	Total        int
	IsRead       bool
	IsBookmarked bool
}

func (m CardModel) Init() tea.Cmd {
	return m.loadStatsCmd()
}

func (m CardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Ensure history is set
	if m.History == nil {
		m.History = history.DefaultStore
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit

		case "n", " ":
			// Next article
			if len(m.Pool) > 0 {
				m.Article = articles.PickRandom(m.Pool)
				// Add to history automatically (as per original logic, though strictly maybe only if they "read" it, but viewed is usually enough)
				if m.Article != nil {
					m.History.AddToHistory(m.Article.URL)
					m.IsRead, m.IsBookmarked = getArticleState(m.History, m.Article.URL)
				}
				m.Message = ""
			} else {
				m.Message = "No more articles in pool!"
				m.MsgType = "error"
			}
			return m, m.loadStatsCmd()

		case "o", "enter":
			// Open
			if m.Article != nil {
				return m, func() tea.Msg {
					if err := OpenBrowser(m.Article.URL); err != nil {
						return statusMsg{text: fmt.Sprintf("Failed to open: %v", err), typeStr: "error"}
					}
					return statusMsg{text: "Opened in browser", typeStr: "success"}
				}
			}

		case "m":
			// Toggle read status
			if m.Article != nil {
				var err error
				msg := "Marked as read"
				typeStr := "success"
				if m.IsRead {
					_, err = m.History.MarkAsUnread(m.Article.URL)
					msg = "Marked as unread"
				} else {
					err = m.History.MarkAsRead(m.Article.URL)
				}
				if err != nil {
					if m.IsRead {
						msg = "Failed to mark as unread"
					} else {
						msg = "Failed to mark as read"
					}
					typeStr = "error"
				} else {
					m.IsRead = !m.IsRead
				}
				m.IsRead, m.IsBookmarked = getArticleState(m.History, m.Article.URL)
				return m, func() tea.Msg {
					return statusMsg{text: msg, typeStr: typeStr}
				}
			}

		case "b":
			// Toggle bookmark
			if m.Article != nil {
				bookmarked, err := m.History.ToggleBookmark(m.Article.URL)
				msg := "Bookmarked"
				typeStr := "info"
				if err != nil {
					msg = "Failed to toggle bookmark"
					typeStr = "error"
				} else if !bookmarked {
					msg = "Removed bookmark"
				} else {
					m.IsBookmarked = true
				}
				if err == nil {
					m.IsBookmarked = bookmarked
				}
				m.IsRead, m.IsBookmarked = getArticleState(m.History, m.Article.URL)
				return m, func() tea.Msg {
					return statusMsg{text: msg, typeStr: typeStr}
				}
			}

		case "h":
			m.ShowHelp = !m.ShowHelp
			return m, nil

		case "y":
			if m.Article != nil {
				err := CopyToClipboard(m.Article.URL)
				msg := "Copied URL to clipboard"
				typeStr := "success"
				if err != nil {
					msg = "Failed to copy URL"
					typeStr = "error"
				}
				return m, func() tea.Msg {
					return statusMsg{text: msg, typeStr: typeStr}
				}
			}
		}

	case statusMsg:
		m.Message = msg.text
		m.MsgType = msg.typeStr
		// clear message after delay?
		cmd := tea.Tick(2*time.Second, func(t time.Time) tea.Msg {
			return clearMsg{}
		})
		if msg.typeStr == "success" && (msg.text == "Marked as read" || msg.text == "Marked as unread") {
			return m, tea.Batch(cmd, m.loadStatsCmd())
		}
		return m, cmd

	case clearMsg:
		m.Message = ""
		return m, nil

	case statsMsg:
		m.Stats = msg.text
		return m, nil
	}

	return m, nil
}

type statusMsg struct {
	text    string
	typeStr string
}

type clearMsg struct{}

type statsMsg struct {
	text string
}

func (m CardModel) View() string {
	if m.Quitting {
		return ""
	}

	if m.Article == nil {
		return "No article selected.\n"
	}

	// Content
	source := articles.FormatSource(m.Article.Source)
	readIcon := UnreadStyle.Render("○")
	if m.IsRead {
		readIcon = ReadStyle.Render("✓")
	}
	bookmarkIcon := " "
	if m.IsBookmarked {
		bookmarkIcon = BookmarkStyle.Render("★")
	}
	content := fmt.Sprintf("%s %s %s\n%s\n\n%s",
		readIcon,
		bookmarkIcon,
		SourceBadgeStyle.Render(source),
		CardTitleStyle.Render(m.Article.Title),
		URLStyle.Render(m.Article.URL),
	)

	card := CardStyle.Render(content)

	statsLine := " "
	if m.Stats != "" {
		statsLine = SubtleStyle.Render(m.Stats)
	}

	// Status line
	status := " "
	if m.Message != "" {
		style := ItemStyle
		if m.MsgType == "success" {
			style = style.Foreground(SuccessColor)
		} else if m.MsgType == "error" {
			style = style.Foreground(ErrorColor)
		} else {
			style = style.Foreground(SecondaryColor)
		}
		status = style.Render(m.Message)
	}

	// Help line
	help := HelpStyle.Render("n: next • o: open • m: mark read • b: bookmark • y: copy URL • h: help • q: quit")
	helpBody := ""
	if m.ShowHelp {
		helpBody = HelpStyle.Render(helpScreen())
	}

	return fmt.Sprintf("\n%s\n%s\n%s\n%s\n%s\n", card, statsLine, status, help, helpBody)
}

func ShowRandomArticle(initial *articles.Article, pool []articles.Article, total int) error {
	return ShowRandomArticleWithStore(initial, pool, total, history.DefaultStore)
}

func ShowRandomArticleWithStore(initial *articles.Article, pool []articles.Article, total int, store history.Store) error {
	// Ensure initial is in history
	if initial != nil {
		store.AddToHistory(initial.URL)
	}
	isRead := false
	isBookmarked := false
	if initial != nil {
		isRead, isBookmarked = getArticleState(store, initial.URL)
	}

	m := CardModel{
		Article:      initial,
		Pool:         pool,
		History:      store,
		Total:        total,
		IsRead:       isRead,
		IsBookmarked: isBookmarked,
	}

	p := tea.NewProgram(m)
	_, err := p.Run()
	return err
}

func (m CardModel) loadStatsCmd() tea.Cmd {
	return func() tea.Msg {
		total := m.Total
		if total == 0 {
			total = len(m.Pool)
				if total == 0 {
					total = len(articles.Cached())
				}
		}
		readUrls, err := m.History.GetReadUrls()
		if err != nil {
			return statsMsg{text: "Stats unavailable"}
		}
		readCount := len(readUrls)
		streak, err := m.History.CalculateStreak()
		if err != nil {
			return statsMsg{text: "Stats unavailable"}
		}
		return statsMsg{text: fmt.Sprintf("Read %d/%d • Streak %d day(s)", readCount, total, streak)}
	}
}

func helpScreen() string {
	return "Shortcuts:\n" +
		"  n / space  next article\n" +
		"  o / enter  open in browser\n" +
		"  m          toggle read\n" +
		"  b          bookmark\n" +
		"  y          copy URL\n" +
		"  h          toggle help\n" +
		"  q          quit\n" +
		"\nLegend:\n" +
		"  ✓ read   ○ unread   ★ bookmarked"
}

func getArticleState(store history.Store, url string) (bool, bool) {
	isRead := false
	isBookmarked := false
	readUrls, err := store.GetReadUrls()
	if err == nil {
		isRead = readUrls[url]
	}
	bookmarks, err := store.GetBookmarks()
	if err == nil {
		for _, entry := range bookmarks {
			if entry.URL == url {
				isBookmarked = true
				break
			}
		}
	}
	return isRead, isBookmarked
}
