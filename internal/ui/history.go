package ui

import (
	"fmt"
	"io"

	"github.com/AfshinJalili/gorandom/internal/articles"
	"github.com/AfshinJalili/gorandom/internal/history"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type HistoryItem struct {
	Entry    history.HistoryEntry
	TitleStr string
}

func (i HistoryItem) Title() string       { return i.TitleStr }
func (i HistoryItem) Description() string { return i.Entry.URL }
func (i HistoryItem) FilterValue() string { return i.TitleStr }

type listKeyMap struct {
	toggleRead     key.Binding
	toggleBookmark key.Binding
	open           key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		toggleRead: key.NewBinding(
			key.WithKeys("m"),
			key.WithHelp("m", "toggle read"),
		),
		toggleBookmark: key.NewBinding(
			key.WithKeys("b"),
			key.WithHelp("b", "bookmark"),
		),
		open: key.NewBinding(
			key.WithKeys("o", "enter"),
			key.WithHelp("enter", "open"),
		),
	}
}

type historyModel struct {
	list         list.Model
	keys         *listKeyMap
	choice       string
	quitting     bool
	msg          string
	historyStore history.Store
}

func (m historyModel) Init() tea.Cmd {
	return nil
}

func (m historyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Ensure store is set
	if m.historyStore == nil {
		m.historyStore = history.DefaultStore
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.toggleRead):
			if i, ok := m.list.SelectedItem().(HistoryItem); ok {
				if i.Entry.IsRead {
					if _, err := m.historyStore.MarkAsUnread(i.Entry.URL); err != nil {
						m.msg = fmt.Sprintf("Failed to mark as unread: %v", err)
						return m, nil
					}
					i.Entry.IsRead = false
					m.msg = "Marked as unread"
				} else {
					if err := m.historyStore.MarkAsRead(i.Entry.URL); err != nil {
						m.msg = fmt.Sprintf("Failed to mark as read: %v", err)
						return m, nil
					}
					i.Entry.IsRead = true
					m.msg = "Marked as read"
				}
				// Refresh item display
				newItem := createListItem(i.Entry)
				idx := m.list.Index()
				m.list.SetItem(idx, newItem)
				return m, nil
			}

		case key.Matches(msg, m.keys.toggleBookmark):
			if i, ok := m.list.SelectedItem().(HistoryItem); ok {
				bookmarked, err := m.historyStore.ToggleBookmark(i.Entry.URL)
				if err != nil {
					m.msg = fmt.Sprintf("Failed to toggle bookmark: %v", err)
					return m, nil
				}
				i.Entry.IsBookmarked = bookmarked
				if bookmarked {
					m.msg = "Bookmarked"
				} else {
					m.msg = "Removed bookmark"
				}
				newItem := createListItem(i.Entry)
				idx := m.list.Index()
				m.list.SetItem(idx, newItem)
				return m, nil
			}

		case key.Matches(msg, m.keys.open):
			if i, ok := m.list.SelectedItem().(HistoryItem); ok {
				if err := OpenBrowser(i.Entry.URL); err != nil {
					m.msg = fmt.Sprintf("Failed to open browser: %v", err)
				} else {
					m.msg = "Opened in browser"
				}
				return m, nil
			}
		}

	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m historyModel) View() string {
	if m.quitting {
		return ""
	}
	return "\n" + m.list.View() + "\n" + HelpStyle.Render(m.msg)
}

// Custom Delegate for rendering
type historyDelegate struct {
	list.DefaultDelegate
}

func (d historyDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	// We rely on the Title() method of the item to return the pre-formatted string.
	// The DefaultDelegate calls Title() and renders it.
}

func newItemDelegate(keys *listKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		return nil
	}

	d.ShortHelpFunc = func() []key.Binding {
		return []key.Binding{keys.open, keys.toggleRead, keys.toggleBookmark}
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{{keys.open, keys.toggleRead, keys.toggleBookmark}}
	}

	return d
}

// Re-implementing simplified HistoryItem to support dynamic title rendering in the main ShowHistory loop
// because pure custom delegate is verbose. We can cheat by formatting the Title string with Lipgloss.

func ShowHistory(entries []history.HistoryEntry, title string) (string, error) {
	return ShowHistoryWithStore(entries, title, history.DefaultStore)
}

func ShowHistoryWithStore(entries []history.HistoryEntry, title string, store history.Store) (string, error) {
	var listItems []list.Item
	for _, entry := range entries {
		listItems = append(listItems, createListItem(entry))
	}

	keys := newListKeyMap()
	delegate := newItemDelegate(keys)

	l := list.New(listItems, delegate, 0, 0)
	l.Title = fmt.Sprintf("%s  (%s)", title, listLegend())
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.Styles.Title = TitleStyle
	l.SetHeight(DefaultListHeight)

	// Disable default help to show our custom help
	l.SetShowHelp(true)

	m := historyModel{list: l, keys: keys, historyStore: store}

	p := tea.NewProgram(m)
	_, err := p.Run()
	if err != nil {
		return "", err
	}
	return "", nil
}

func createListItem(entry history.HistoryEntry) HistoryItem {
	var title string
	for _, a := range articles.Data {
		if a.URL == entry.URL {
			title = a.Title
			break
		}
	}
	if title == "" {
		title = "Unknown"
	}

	// Icons
	readIcon := UnreadStyle.Render("○")
	if entry.IsRead {
		readIcon = ReadStyle.Render("✓")
	}

	bookmarkIcon := " "
	if entry.IsBookmarked {
		bookmarkIcon = BookmarkStyle.Render("★")
	}

	date := entry.ViewedAt.Format("2006-01-02")

	// Construct the display title
	// "✓ ★ 2025-01-01 - Title"
	fullTitle := fmt.Sprintf("%s %s %s - %s", readIcon, bookmarkIcon, SubtleStyle.Render(date), title)

	return HistoryItem{
		Entry:    entry,
		TitleStr: fullTitle,
	}
}
