package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type teaProgram interface {
	Run() (tea.Model, error)
}

var newTeaProgram = func(m tea.Model, opts ...tea.ProgramOption) teaProgram {
	return tea.NewProgram(m, opts...)
}

var withAltScreen = tea.WithAltScreen

// RunProgram standardizes TUI execution with alternate screen enabled.
func RunProgram(m tea.Model) (tea.Model, error) {
	p := newTeaProgram(m, withAltScreen())
	return p.Run()
}
