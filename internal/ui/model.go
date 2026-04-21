package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/sections"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

// Model is the root Bubble Tea model
type Model struct {
	welcome sections.Welcome
}

// NewModel returns an initialized root Model
func NewModel(theme styles.Theme) Model {
	return Model{
		welcome: sections.NewWelcome(theme, 42),
	}
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
		}
	}
	return m, nil
}

// View implements tea.Model
func (m Model) View() string {
	return m.welcome.View()
}
