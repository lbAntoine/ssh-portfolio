package ui_test

import (
	"testing"

	tea "charm.land/bubbletea/v2"
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
	for range len(splash.SplashText) + len(splash.SplashSubtitle) + 30 {
		m, _ = m.Update(splash.TickMsg{})
	}
	if !m.(ui.App).Ready() {
		t.Error("app should be ready after splash completes")
	}
}

func TestApp_KeypressSkipsSplash(t *testing.T) {
	a := ui.NewApp(styles.Minimal(), 0)
	m, _ := a.Update(tea.KeyPressMsg{Code: 'x', Text: "x"})
	// second keypress after skip should transition
	m, _ = m.Update(tea.KeyPressMsg{Code: 'x', Text: "x"})
	if !m.(ui.App).Ready() {
		t.Error("app should be ready after skipping splash")
	}
}

func TestApp_Init(t *testing.T) {
	a := ui.NewApp(styles.Minimal(), 0)
	_ = a.Init() // delegates to splash.Init, just verify no panic
}

func TestApp_View_SplashState(t *testing.T) {
	a := ui.NewApp(styles.Minimal(), 0)
	v := a.View()
	if v.Content == "" {
		t.Error("expected non-empty view in splash state")
	}
	if !v.AltScreen {
		t.Error("expected AltScreen true in splash state")
	}
}

func TestApp_View_ReadyState(t *testing.T) {
	a := ui.NewApp(styles.Minimal(), 0)
	var m tea.Model = a
	for range len(splash.SplashText) + len(splash.SplashSubtitle) + 30 {
		m, _ = m.Update(splash.TickMsg{})
	}
	v := m.(ui.App).View()
	if v.Content == "" {
		t.Error("expected non-empty view in ready state")
	}
	if !v.AltScreen {
		t.Error("expected AltScreen true in ready state")
	}
}

func TestApp_WindowSizeMsg(t *testing.T) {
	a := ui.NewApp(styles.Minimal(), 0)
	next, _ := a.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	if next == nil {
		t.Error("expected non-nil model after WindowSizeMsg")
	}
}
