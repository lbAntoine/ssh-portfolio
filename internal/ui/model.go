package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/sections"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

const (
	maxContentWidth  = 96
	maxContentHeight = 36
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
	width       int
	height      int
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

func (m Model) contentSize() (w, h int) {
	w = m.width - 4
	if w > maxContentWidth {
		w = maxContentWidth
	}
	h = m.height - 4
	if h > maxContentHeight {
		h = maxContentHeight
	}
	return
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
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		cw, ch := m.contentSize()
		for i, s := range m.sections {
			if sz, ok := s.(interface{ SetSize(int, int) }); ok {
				sz.SetSize(cw, ch)
				m.sections[i] = s
			}
		}
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab", "l":
			m.active = (m.active + 1) % len(m.sections)
			return m, nil
		case "shift+tab", "h":
			m.active = (m.active - 1 + len(m.sections)) % len(m.sections)
			return m, nil
		case "?":
			m.helpVisible = !m.helpVisible
			return m, nil
		default:
			if len(msg.Text) == 1 && msg.Text[0] >= '1' && msg.Text[0] <= '7' {
				m.active = int(msg.Text[0] - '1')
				return m, nil
			}
		}
	}

	updated, cmd := m.sections[m.active].Update(msg)
	m.sections[m.active] = updated
	return m, cmd
}

// View implements tea.Model
func (m Model) View() tea.View {
	cw, ch := m.contentSize()

	var content string
	if m.helpVisible {
		content = m.helpView()
	} else {
		content = m.sections[m.active].View().Content
	}

	tabBar := lipgloss.NewStyle().Width(cw).Render(m.tabBar())
	box := lipgloss.Place(cw, ch, lipgloss.Left, lipgloss.Center, content)
	block := lipgloss.JoinVertical(lipgloss.Left, tabBar, "", box)

	if m.width == 0 || m.height == 0 {
		return tea.NewView(block)
	}
	return tea.NewView(lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, block))
}

func (m Model) tabBar() string {
	if styles.BreakpointFor(m.width) == styles.Compact {
		return m.theme.Muted.Render(
			fmt.Sprintf("[%d/7] %s", m.active+1, sectionNames[m.active]),
		)
	}
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
