package ui

import "github.com/charmbracelet/lipgloss"

const (
	DefaultListHeight = 20
	StatLabelWidth    = 20
	ProgressBarWidth  = 20
)

var (
	// Colors
	PrimaryColor   = lipgloss.Color("#7D56F4")
	SecondaryColor = lipgloss.Color("#EE6FF8")
	SubtleColor    = lipgloss.Color("#626262")
	HighlightColor = lipgloss.Color("#FAFAFA")
	ErrorColor     = lipgloss.Color("#FF5F87")
	SuccessColor   = lipgloss.Color("#42E66C")
	WarningColor   = lipgloss.Color("#F3F99D") // Yellowish

	// Styles
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(HighlightColor).
			Background(PrimaryColor).
			Padding(0, 1)

	ItemStyle = lipgloss.NewStyle().
			PaddingLeft(2) // Reduced padding as we have icons now

	SelectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(0). // Delegate handles indentation
				Foreground(SecondaryColor)

	PaginationStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(SubtleColor)

	// Card Styles
	CardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(PrimaryColor).
			Padding(1, 2)

	CardTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(SecondaryColor).
			MarginBottom(1)

	SourceBadgeStyle = lipgloss.NewStyle().
				Foreground(HighlightColor).
				Background(SubtleColor).
				Padding(0, 1).
				MarginRight(1)

	URLStyle = lipgloss.NewStyle().
			Foreground(SubtleColor).
			Underline(true)

	HelpStyle = lipgloss.NewStyle().
			Foreground(SubtleColor).
			MarginTop(1)

	// Stats Styles
	StatsStyle = lipgloss.NewStyle().
			Foreground(HighlightColor).
			Bold(true)

	StatLabelStyle = lipgloss.NewStyle().
			Foreground(SubtleColor).
			Width(StatLabelWidth)

	StatValueStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(HighlightColor)

	StatBarStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor)

	SubtleStyle = lipgloss.NewStyle().
			Foreground(SubtleColor)

	// List Indicators
	ReadStyle = lipgloss.NewStyle().
			Foreground(SuccessColor).
			Bold(true)

	UnreadStyle = lipgloss.NewStyle().
			Foreground(SubtleColor)

	BookmarkStyle = lipgloss.NewStyle().
			Foreground(WarningColor)
)
