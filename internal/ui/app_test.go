package ui_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lbAntoine/ssh-portfolio/internal/ui"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/splash"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

func TestApp_StartsWithSplash(t *testing.T) {
	a := ui.NewApp(styles.Minimal(), 0)
	if a.Ready() {
		t.Error("app should start in splash state, not ready")
	}
}

func TestApp_TransitionsAfterSplashDone(t *testing.T) {
	a := ui.NewApp(styles.Minimal(), 0)
	// Tick through all splash chars
	var m tea.Model = a
	for range len(splash.SplashText) + 12 {
		m, _ = m.Update(splash.TickMsg{})
	}
	if !m.(ui.App).Ready() {
		t.Error("app should be ready after splash completes")
	}
}

func TestApp_KeypressSkipsSplash(t *testing.T) {
	a := ui.NewApp(styles.Minimal(), 0)
	m, _ := a.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("x")})
	// second keypress after skip should transition
	m, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("x")})
	if !m.(ui.App).Ready() {
		t.Error("app should be ready after skipping splash")
	}
}
