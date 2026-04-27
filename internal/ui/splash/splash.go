package splash

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

// SplashText and SplashSubtitle are the default content for the typewriter animation.
const (
	SplashText     = "antoine le bras"
	SplashSubtitle = "backend developer"
	tickInterval   = 80 * time.Millisecond
	pauseTicks     = 6
	holdTicks      = 12
)

// phase constants for the two-phase typewriter animation
const (
	phaseMain     = iota // revealing main text
	phasePause           // brief pause between main and subtitle
	phaseSubtitle        // revealing subtitle
	phaseHold            // holding before transition
)

// TickMsg is sent on each tick interval
type TickMsg struct{}

// Model is the typewriter splash screen
type Model struct {
	text        string
	subtitle    string
	theme       styles.Theme
	phase       int
	revealed    int
	revealedSub int
	cursor      bool
	ticks       int
	done        bool
	width       int
	height      int
}

// NewSplash returns an initialized splash model
func NewSplash(text, subtitle string, theme styles.Theme) Model {
	return Model{text: text, subtitle: subtitle, theme: theme, cursor: true}
}

// Done returns true when the animation has completed
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
		switch m.phase {
		case phaseMain:
			if m.revealed < len(m.text) {
				m.revealed++
				m.cursor = true
				return m, tick()
			}
			m.phase = phasePause
			m.ticks = 0
			return m, tick()

		case phasePause:
			m.ticks++
			m.cursor = m.ticks%2 == 0
			if m.ticks >= pauseTicks {
				m.phase = phaseSubtitle
				m.ticks = 0
				m.cursor = true
			}
			return m, tick()

		case phaseSubtitle:
			if m.revealedSub < len(m.subtitle) {
				m.revealedSub++
				m.cursor = true
				return m, tick()
			}
			m.phase = phaseHold
			m.ticks = 0
			return m, tick()

		case phaseHold:
			m.ticks++
			m.cursor = m.ticks%2 == 0
			if m.ticks >= holdTicks {
				m.done = true
				return m, nil
			}
			return m, tick()
		}

	case tea.KeyPressMsg:
		m.revealed = len(m.text)
		m.revealedSub = len(m.subtitle)
		m.done = true
	}
	return m, nil
}

// View implements tea.Model
func (m Model) View() tea.View {
	t := m.theme

	mainText := t.Accent.Render(m.text[:m.revealed])
	if m.phase == phaseMain && !m.done {
		if m.cursor {
			mainText += t.Muted.Render("▋")
		}
	}

	var rendered string
	if m.phase >= phaseSubtitle || m.done {
		sub := t.Muted.Render(m.subtitle[:m.revealedSub])
		if m.phase == phaseSubtitle && !m.done {
			if m.cursor {
				sub += t.Muted.Render("▋")
			}
		}
		rendered = mainText + "\n" + sub
	} else {
		rendered = mainText
	}

	if m.width == 0 || m.height == 0 {
		return tea.NewView(rendered)
	}
	return tea.NewView(lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, rendered))
}
