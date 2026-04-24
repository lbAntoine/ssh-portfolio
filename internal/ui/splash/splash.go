package splash

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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
	text       string
	theme      styles.Theme
	revealed   int
	cursor     bool
	held       int
	blinkCount int
	done       bool
	width      int
	height     int
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
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case TickMsg:
		if m.revealed < len(m.text) {
			m.revealed++
			m.cursor = true
			return m, tick()
		}
		m.held++
		m.blinkCount++
		if m.blinkCount%4 == 0 {
			m.cursor = !m.cursor
		}
		if m.held >= holdTicks {
			m.done = true
			return m, nil
		}
		return m, tick()
	case tea.KeyPressMsg:
		m.revealed = len(m.text)
		m.done = true
	}
	return m, nil
}

// View implements tea.Model
func (m Model) View() tea.View {
	text := m.theme.Accent.Render(m.text[:m.revealed])
	if !m.done {
		cursor := ""
		if m.cursor {
			cursor = m.theme.Muted.Render("▋")
		}
		text = text + cursor
	}
	if m.width == 0 || m.height == 0 {
		return tea.NewView(text)
	}
	return tea.NewView(lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, text))
}
