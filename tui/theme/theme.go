package theme

import (
	"charm.land/lipgloss/v2"
)

var (
	PrimaryColor   = lipgloss.Color("#83C092")
	SecondaryColor = lipgloss.Color("#98AF77")
	ErrorColor     = lipgloss.Color("#FF4C4C")
	MutedColor     = lipgloss.Color("#626262")
	BgAccentColor  = lipgloss.Color("#1E1E2E")
	MatchHighlight = lipgloss.Color("#F38BA8")

	BaseBox = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(PrimaryColor).
		Padding(0, 1)

	FocusedBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(PrimaryColor).
			Padding(0, 1)

	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(PrimaryColor).
			Padding(0, 1)

	MatchBadge = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(SecondaryColor).
			Padding(0, 1)

	ErrorBadge = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(ErrorColor).
			Padding(0, 1)
)
