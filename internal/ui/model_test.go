package ui_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/lbAntoine/ssh-portfolio/internal/ui"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

func TestRootModel_InitialView(t *testing.T) {
	m := ui.NewModel(styles.Minimal())
	view := m.View()
	if view == "" {
		t.Error("expected non-empty view on init")
	}
}

func TestRootModel_QuitOnQ(t *testing.T) {
	m := ui.NewModel(styles.Minimal())
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")})
	if cmd == nil {
		t.Fatal("expected a command, got nil")
	}
	if cmd() != tea.Quit() {
		t.Error("expected tea.Quit command on 'q'")
	}
}

func TestRootModel_QuitOnCtrlC(t *testing.T) {
	m := ui.NewModel(styles.Minimal())
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	if cmd == nil {
		t.Fatal("expected a command, got nil")
	}
	if cmd() != tea.Quit() {
		t.Error("expected tea.Quit command on ctrl+c")
	}
}
