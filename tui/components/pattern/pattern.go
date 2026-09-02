package pattern

import (
	"fmt"

	"regex/tui/engine"
	"regex/tui/theme"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// Model holds state for getting input of data modularly
// It consists for actual input, any errors, whether it's focused or not
// and finally, it's width
type Model struct {
	input       textinput.Model
	syntaxError *engine.SyntaxError
	focused     bool
	width       int
}

func New() Model {
	ti := textinput.New()
	ti.Placeholder = "Some kinda Regex prolly"
	ti.Focus()

	return Model{
		input:   ti,
		focused: true,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m *Model) SetError(err *engine.SyntaxError) {
	m.syntaxError = err
}

func (m *Model) SetWidth(w int) {
	m.width = w
	m.input.SetWidth(m.width)
}

func (m *Model) Focus() tea.Cmd {
	m.focused = true
	return m.input.Focus()
}

func (m *Model) Blur() {
	m.focused = false
	m.input.Blur()
}

func (m Model) Value() string {
	return m.input.Value()
}

func (m Model) View() string {
	box := theme.BaseBox
	if m.focused {
		box = theme.FocusedBox
	}
	box = box.Width(m.width)

	header := "REGEX PATTERN"
	if m.syntaxError != nil {
		header = fmt.Sprintf("REGEX PATTER %s", theme.ErrorBadge.Render(m.syntaxError.Message))
	}

	content := lipgloss.JoinVertical(lipgloss.Left, header, m.input.View())
	return box.Render(content)
}
