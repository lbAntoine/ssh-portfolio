package splash

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

// SplashText is the text displayed by the splash component
const (
	SplashText   = "antoine le bras"
	tickInterval = 80 * time.Millisecond
	holdTicks    = 12
)

// TickMsg is sent on each tick interval
type TickMsg struct{}

// Model is the typewriter splash screen
type Model struct {
	text     string
	theme    styles.Theme
	revealed int
	cursor   bool
	held     int
	done     bool
}

// NewSplash returns an initialized splash model
func NewSplash(text string, theme styles.Theme) Model {
	return Model{text: text, theme: theme, cursor: true}
}

// Done returns true when all chars have been revealed
func (m Model) Done() bool { return m.done }

func tick() tea.Cmd {
	return tea.Tick(tickInterval, func(time.Time) tea.Msg { return TickMsg{} })
}

// Init implements tea.Model
func (m Model) Init() tea.Cmd { return tick() }

// Update implements tea.Model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case TickMsg:
		if m.revealed < len(m.text) {
			m.revealed++
			m.cursor = true
			return m, tick()
		}
		m.cursor = !m.cursor
		m.held++
		if m.held >= holdTicks {
			m.done = true
			return m, nil
		}
		return m, tick()
	case tea.KeyMsg:
		m.revealed = len(m.text)
		m.done = true
	}
	return m, nil
}

// View implements tea.Model
func (m Model) View() string {
	text := m.theme.Accent.Render(m.text[:m.revealed])
	if m.done {
		return text
	}
	cursor := ""
	if m.cursor {
		cursor = m.theme.Muted.Render("▋")
	}
	return text + cursor
}
