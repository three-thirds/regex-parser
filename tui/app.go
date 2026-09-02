// Package tui consists of critical methods written for main AppModel
package tui

import (
	"fmt"
	"strings"

	"regex/tui/components/pattern"
	"regex/tui/engine"
	"regex/tui/theme"

	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type activePane int

// Simple enum for Panes
const (
	panePattern activePane = iota
	paneTestString
)

// AppModel is main app state which holds essential information
// like matches from Parser engine, currently selecter pane, input string,
// pattern to compile, previous result, etc.
type AppModel struct {
	engine     engine.Engine
	activePane activePane

	patternComp pattern.Model
	testString  textarea.Model

	lastResult engine.AnalysisResult
	width      int
	height     int
}

// NewApp initializes a new AppModel state struct for application
func NewApp(eng engine.Engine) AppModel {
	ta := textarea.New()
	ta.Placeholder = "Insert test string to match against...."

	return AppModel{
		engine:      eng,
		activePane:  panePattern,
		patternComp: pattern.New(),
		testString:  ta,
	}
}

//=========== Interface Methods to be applied so that AppModel is qualified to be passed as tea.Model ===========//

func (m AppModel) Init() tea.Cmd {
	return m.patternComp.Init()
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.recalculateSizes()

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "tab":
			if m.activePane == panePattern {
				m.activePane = paneTestString
				m.patternComp.Blur()
				cmds = append(cmds, m.testString.Focus())
			} else {
				m.activePane = panePattern
				m.testString.Blur()
				cmds = append(cmds, m.patternComp.Focus())
			}
			return m, tea.Batch(cmds...)
		}
	}

	if m.activePane == panePattern {
		var cmd tea.Cmd
		m.patternComp, cmd = m.patternComp.Update(msg)
		cmds = append(cmds, cmd)
	} else {

		var cmd tea.Cmd
		m.testString, cmd = m.testString.Update(msg)
		cmds = append(cmds, cmd)
	}

	m.lastResult = m.engine.Evaluate(m.patternComp.Value(), m.testString.Value())
	m.patternComp.SetError(m.lastResult.SyntaxError)

	return m, tea.Batch(cmds...)
}

func (m AppModel) View() tea.View {
	if m.width == 0 {
		return tea.NewView("Initializing...")
	}

	halfWidth := (m.width / 2) - 2

	testBox := theme.BaseBox
	if m.activePane == paneTestString {
		testBox = theme.FocusedBox
	}

	testView := testBox.Width(halfWidth).Render(
		lipgloss.JoinVertical(lipgloss.Left, "TEST STRING", m.testString.View()),
	)

	leftColumn := lipgloss.JoinVertical(lipgloss.Left, m.patternComp.View(), testView)

	var matchDetails strings.Builder
	if len(m.lastResult.Matches) > 0 {
		matchDetails.WriteString(theme.MatchBadge.Render(fmt.Sprintf("%d Match(es) Found", len(m.lastResult.Matches))) + "\n\n")
		for _, match := range m.lastResult.Matches {
			matchDetails.WriteString(fmt.Sprintf("• Match [%d-%d]: %q\n", match.Start, match.End, match.Value))
			for _, grp := range match.Groups {
				matchDetails.WriteString(fmt.Sprintf("   └ Group %d (%s): %q\n", grp.Index, grp.Name, grp.Value))
			}
		}
	} else {
		matchDetails.WriteString(lipgloss.NewStyle().Foreground(theme.MutedColor).Render("No matches."))
	}

	matchBox := theme.BaseBox.Width(halfWidth).
		Render(lipgloss.JoinVertical(lipgloss.Left, "MATCH INFORMATION", matchDetails.String()))

	astBox := theme.BaseBox.Width(halfWidth).Render(
		lipgloss.JoinVertical(lipgloss.Left, "AST EXPLANATION", m.lastResult.ASTFormatted),
	)

	rightColumn := lipgloss.JoinVertical(lipgloss.Left, matchBox, astBox)

	title := theme.TitleStyle.Render("Regex101 TUI")
	body := lipgloss.JoinHorizontal(lipgloss.Top, leftColumn, rightColumn)
	status := lipgloss.NewStyle().Foreground(theme.MutedColor).Render(
		fmt.Sprintf(" [Tab] Switch Pane  •  [Ctrl+C] Quit  •  Eval Time: %s", m.lastResult.ExecutionTime),
	)

	fullContent := lipgloss.JoinVertical(lipgloss.Left, title, body, status)

	view := tea.NewView(fullContent)
	view.AltScreen = true // Keeps your terminal clean by rendering in alt screen mode!
	return view
}

// == Helper functions ==//

// recalculateSizes caluculates the new UI size everytime
// a new WindowSizeMsg is received
func (m *AppModel) recalculateSizes() {
	halfWidth := (m.width / 2) - 2
	m.patternComp.SetWidth(halfWidth)
	m.testString.SetWidth(halfWidth - 4)
	m.testString.SetHeight(m.height - 14)
}
