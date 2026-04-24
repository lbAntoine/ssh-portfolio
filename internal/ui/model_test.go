package ui_test

import (
	"testing"

	tea "charm.land/bubbletea/v2"

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
	_, cmd := m.Update(tea.KeyPressMsg{Code: 'q', Text: "q"})
	if cmd == nil {
		t.Fatal("expected a command, got nil")
	}
	if cmd() != tea.Quit() {
		t.Error("expected tea.Quit command on 'q'")
	}
}

func TestRootModel_QuitOnCtrlC(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	_, cmd := m.Update(tea.KeyPressMsg{Code: 'c', Mod: tea.ModCtrl})
	if cmd == nil {
		t.Fatal("expected a command, got nil")
	}
	if cmd() != tea.Quit() {
		t.Error("expected tea.Quit command on ctrl+c")
	}
}

func TestRoot_TabAdvancesSection(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	next, _ := m.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	if next.(ui.Model).ActiveSection() != 1 {
		t.Errorf("expected section 1 after tab, got %d", next.(ui.Model).ActiveSection())
	}
}

func TestRoot_ShiftTabGoesBack(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	next, _ := m.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	next, _ = next.(ui.Model).Update(tea.KeyPressMsg{Code: tea.KeyTab, Mod: tea.ModShift})
	if next.(ui.Model).ActiveSection() != 0 {
		t.Errorf("expected section 0 after shift+tab, got %d", next.(ui.Model).ActiveSection())
	}
}

func TestRoot_VimNextSection(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	next, _ := m.Update(tea.KeyPressMsg{Code: 'l', Text: "l"})
	if next.(ui.Model).ActiveSection() != 1 {
		t.Errorf("expected section 1 after vim 'l' move, got %d", next.(ui.Model).ActiveSection())
	}
}

func TestRoot_VimPrevSection(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	next, _ := m.Update(tea.KeyPressMsg{Code: 'l', Text: "l"})
	next, _ = next.(ui.Model).Update(tea.KeyPressMsg{Code: 'h', Text: "h"})
	if next.(ui.Model).ActiveSection() != 0 {
		t.Errorf("expected section 0 after vim 'h' move, got %d", next.(ui.Model).ActiveSection())
	}
}

func TestRoot_NumberKeyJumpsDirect(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	next, _ := m.Update(tea.KeyPressMsg{Code: '3', Text: "3"})
	if next.(ui.Model).ActiveSection() != 2 {
		t.Errorf("expected section 2 on key '3', got %d", next.(ui.Model).ActiveSection())
	}
}

func TestRoot_WrapAroundAtEnd(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	next, _ := m.Update(tea.KeyPressMsg{Code: '7', Text: "7"})
	next, _ = next.(ui.Model).Update(tea.KeyPressMsg{Code: tea.KeyTab})
	if next.(ui.Model).ActiveSection() != 0 {
		t.Errorf("expected wrap to section 0, got %d", next.(ui.Model).ActiveSection())
	}
}

func TestRoot_WrapAroundAtStart(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	next, _ := m.Update(tea.KeyPressMsg{Code: tea.KeyTab, Mod: tea.ModShift})
	if next.(ui.Model).ActiveSection() != 6 {
		t.Errorf("expected wrap to section 6, got %d", next.(ui.Model).ActiveSection())
	}
}

func TestRoot_HelpToggleOnQuestionMark(t *testing.T) {
	m := ui.NewModel(styles.Minimal(), 42)
	next, _ := m.Update(tea.KeyPressMsg{Code: '?', Text: "?"})
	if !next.(ui.Model).HelpVisible() {
		t.Error("expected help to be visible after '?'")
	}
	next, _ = next.(ui.Model).Update(tea.KeyPressMsg{Code: '?', Text: "?"})
	if next.(ui.Model).HelpVisible() {
		t.Error("expected help to be hidden after second '?'")
	}
}
