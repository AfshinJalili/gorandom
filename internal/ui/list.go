package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	title string
	desc  string
	value string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i.value
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" || m.quitting {
		return ""
	}
	return "\n" + m.list.View() + "\n" + HelpStyle.Render(listLegend())
}

func SelectArticle(items []struct{ Title, Value string }, title string) (string, error) {
	var listItems []list.Item
	for _, i := range items {
		listItems = append(listItems, item{title: i.Title, value: i.Value, desc: i.Value})
	}

	l := list.New(listItems, list.NewDefaultDelegate(), 0, 0)
	l.Title = fmt.Sprintf("%s  (%s)", title, listLegend())
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	// simple styles
	l.Styles.Title = TitleStyle

	// Approx height
	l.SetHeight(20)

	m := model{list: l}

	p := tea.NewProgram(m)
	finalModel, err := p.Run()
	if err != nil {
		return "", err
	}

	finalM, ok := finalModel.(model)
	if !ok {
		return "", fmt.Errorf("internal error: could not assert model")
	}

	if finalM.quitting {
		return "", nil
	}

	return finalM.choice, nil
}

func listLegend() string {
	return strings.Join([]string{
		"✓ read",
		"○ unread",
		"★ bookmarked",
	}, " • ")
}
