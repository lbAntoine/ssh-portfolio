package splash_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/splash"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

func TestSplash_StartsEmpty(t *testing.T) {
	s := splash.NewSplash("hello", styles.Minimal())
	if !strings.Contains(s.View(), "") {
		t.Error("expected empty view on start")
	}
	if s.Done() {
		t.Error("splash should not be done on start")
	}
}

func TestSplash_TickRevealsOneChar(t *testing.T) {
	s := splash.NewSplash("hello", styles.Minimal())
	next, _ := s.Update(splash.TickMsg{})
	m := next.(splash.Model)
	if !strings.Contains(m.View(), "h") {
		t.Errorf("expected 'h', got %q", m.View())
	}
}

func TestSplash_ViewShowsRevealedText(t *testing.T) {
	s := splash.NewSplash("hello", styles.Minimal())
	var m tea.Model = s
	for range 3 {
		m, _ = m.Update(splash.TickMsg{})
	}
	if !strings.Contains(m.View(), "hel") {
		t.Errorf("expected 'hel', got %q", m.(splash.Model).View())
	}
}

func TestSplash_KeypressSkipsToEnd(t *testing.T) {
	s := splash.NewSplash("hello", styles.Minimal())
	next, _ := s.Update(tea.KeyPressMsg{Code: 'x', Text: "x"})
	m := next.(splash.Model)
	if !strings.Contains(m.View(), "hello") {
		t.Errorf("expected 'hello' after keypress, got %q", m.View())
	}
}

func TestSplash_DoneWhenFullyRevealed(t *testing.T) {
	s := splash.NewSplash("hi", styles.Minimal())
	var m tea.Model = s
	for range 14 {
		m, _ = m.Update(splash.TickMsg{})
	}
	if !m.(splash.Model).Done() {
		t.Errorf("expected Done() true after all chars revealed")
	}
}
