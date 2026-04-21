package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Model is the root Bubble Tea model
type Model struct{}

func NewModel() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

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

func (m Model) View() string {
	return "hello -- ssh-portfolio\n\npress q to quit"
}
