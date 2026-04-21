package ui_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/lbAntoine/ssh-portfolio/internal/ui"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

func TestRootModel_InitialView(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	view := m.View()
	if view == "" {
		t.Error("expected non-empty view on init")
	}
}

func TestRootModel_QuitOnQ(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")})
	if cmd == nil {
		t.Fatal("expected a command, got nil")
	}
	if cmd() != tea.Quit() {
		t.Error("expected tea.Quit command on 'q'")
	}
}

func TestRootModel_QuitOnCtrlC(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	if cmd == nil {
		t.Fatal("expected a command, got nil")
	}
	if cmd() != tea.Quit() {
		t.Error("expected tea.Quit command on ctrl+c")
	}
}

func TestRoot_TabAdvancesSection(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyTab})
	if next.(ui.Model).ActiveSection() != 1 {
		t.Errorf("expected section 1 after tab, got %d", next.(ui.Model).ActiveSection())
	}
}

func TestRoot_ShiftTabGoesBack(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyTab})
	next, _ = next.(ui.Model).Update(tea.KeyMsg{Type: tea.KeyShiftTab})
	if next.(ui.Model).ActiveSection() != 0 {
		t.Errorf("expected section 0 after shift+tab, got %d", next.(ui.Model).ActiveSection())
	}
}

func TestRoot_NumberKeyJumpsDirect(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("3")})
	if next.(ui.Model).ActiveSection() != 2 {
		t.Errorf("expected section 2 on key '3', got %d", next.(ui.Model).ActiveSection())
	}
}

func TestRoot_WrapAroundAtEnd(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("7")})
	next, _ = next.(ui.Model).Update(tea.KeyMsg{Type: tea.KeyTab})
	if next.(ui.Model).ActiveSection() != 0 {
		t.Errorf("expected wrap to section 0, got %d", next.(ui.Model).ActiveSection())
	}
}

func TestRoot_WrapAroundAtStart(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyShiftTab})
	if next.(ui.Model).ActiveSection() != 6 {
		t.Errorf("expected wrap to section 6, got %d", next.(ui.Model).ActiveSection())
	}
}

func TestRoot_HelpToggleOnQuestionMark(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("?")})
	if !next.(ui.Model).HelpVisible() {
		t.Error("expected help to be visible after '?'")
	}
	next, _ = next.(ui.Model).Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("?")})
	if next.(ui.Model).HelpVisible() {
		t.Error("expected help to be hidden after second '?'")
	}
}
