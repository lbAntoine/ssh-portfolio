package sections_test

import (
	"strings"
	"testing"

	"github.com/lbAntoine/ssh-portfolio/internal/ui/sections"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

func TestWelcome_RendersName(t *testing.T) {
	w := sections.NewWelcome(styles.Minimal(), 42)
	if !strings.Contains(w.View().Content, "Antoine") {
		t.Error("expected name 'Antoine' in welcome view")
	}
}

func TestWelcome_RendersKoreanGreeting(t *testing.T) {
	w := sections.NewWelcome(styles.Minimal(), 42)
	w.SetSize(100, 40)
	if !strings.Contains(w.View().Content, "안녕하세요") {
		t.Error("expected Korean greeting in welcome view")
	}
}

func TestWelcome_CompactHidesGreeting(t *testing.T) {
	w := sections.NewWelcome(styles.Minimal(), 42)
	w.SetSize(50, 20)
	if strings.Contains(w.View().Content, "안녕하세요") {
		t.Error("compact mode should hide Korean greeting")
	}
}

func TestWelcome_ShowsVisitorCount(t *testing.T) {
	w := sections.NewWelcome(styles.Minimal(), 42)
	if !strings.Contains(w.View().Content, "42") {
		t.Error("expected visitor count 42 in welcome view")
	}
}
