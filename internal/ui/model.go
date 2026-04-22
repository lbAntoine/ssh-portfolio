package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/sections"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

var sectionNames = []string{
	"welcome", "about", "projects", "stack", "now", "contact", "resume",
}

// Model is the root Bubble Tea model
type Model struct {
	theme       styles.Theme
	sections    []tea.Model
	active      int
	helpVisible bool
}

// NewModel returns an initialized root Model
func NewModel(theme styles.Theme, visitorCount int) Model {
	return Model{
		theme: theme,
		sections: []tea.Model{
			sections.NewWelcome(theme, visitorCount),
			sections.NewAbout(theme),
			sections.NewProject(theme),
			sections.NewStack(theme),
			sections.NewNow(theme),
			sections.NewContact(theme),
			sections.NewResume(theme),
		},
	}
}

// ActiveSection returns the index of the currently active section
func (m Model) ActiveSection() int {
	return m.active
}

// HelpVisible return whether the help overlay is shown
func (m Model) HelpVisible() bool {
	return m.helpVisible
}

// Init implements tea.Model
func (m Model) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case msg.Type == tea.KeyCtrlC,
			msg.Type == tea.KeyRunes && string(msg.Runes) == "q":
			return m, tea.Quit

		case msg.Type == tea.KeyTab,
			msg.Type == tea.KeyRunes && string(msg.Runes) == "l":
			m.active = (m.active + 1) % len(m.sections)
			return m, nil

		case msg.Type == tea.KeyShiftTab,
			msg.Type == tea.KeyRunes && string(msg.Runes) == "h":
			m.active = (m.active - 1 + len(m.sections)) % len(m.sections)
			return m, nil

		case msg.Type == tea.KeyRunes && string(msg.Runes) == "?":
			m.helpVisible = !m.helpVisible
			return m, nil

		case msg.Type == tea.KeyRunes:
			r := string(msg.Runes)
			if len(r) == 1 && r[0] >= '1' && r[0] <= '7' {
				m.active = int(r[0] - '1')
				return m, nil
			}
		}
	}

	updated, cmd := m.sections[m.active].Update(msg)
	m.sections[m.active] = updated
	return m, cmd
}

// View implements tea.Model
func (m Model) View() string {
	if m.helpVisible {
		return m.helpView()
	}
	return m.tabBar() + "\n\n" + m.sections[m.active].View()
}

func (m Model) tabBar() string {
	var tabs []string
	for i, name := range sectionNames {
		if i == m.active {
			tabs = append(tabs, m.theme.Accent.Bold(true).Render(fmt.Sprintf("[%s]", name)))
		} else {
			tabs = append(tabs, m.theme.Muted.Render(fmt.Sprintf(" %s ", name)))
		}
	}
	return strings.Join(tabs, m.theme.Muted.Render("·"))
}

func (m Model) helpView() string {
	help := lipgloss.NewStyle().Padding(1, 2).Render(
		m.theme.Title.Render("keybindings") + "\n\n" +
			m.theme.Body.Render("tab / l        ") + "  " + m.theme.Muted.Render("next section") + "\n" +
			m.theme.Body.Render("shift+tab / h  ") + "  " + m.theme.Muted.Render("prev section") + "\n" +
			m.theme.Body.Render("1–7            ") + "  " + m.theme.Muted.Render("jump to section") + "\n" +
			m.theme.Body.Render("?              ") + "  " + m.theme.Muted.Render("toggle help") + "\n" +
			m.theme.Body.Render("q / ctrl+c     ") + "  " + m.theme.Muted.Render("quit"),
	)
	return help
}
