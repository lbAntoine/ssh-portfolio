package sections_test

import (
	"strings"
	"testing"

	"github.com/lbAntoine/ssh-portfolio/internal/ui/sections"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

func TestAbout_RendersContent(t *testing.T) {
	s := sections.NewAbout(styles.Minimal())
	if strings.TrimSpace(s.View()) == "" {
		t.Error("expected non-empty about view")
	}
}

func TestProjects_RendersContent(t *testing.T) {
	s := sections.NewProject(styles.Minimal())
	if strings.TrimSpace(s.View()) == "" {
		t.Error("expected non-empty projects view")
	}
}

func TestStack_RendersContent(t *testing.T) {
	s := sections.NewStack(styles.Minimal())
	if strings.TrimSpace(s.View()) == "" {
		t.Error("expected non-empty stack view")
	}
}

func TestNow_RendersContent(t *testing.T) {
	s := sections.NewNow(styles.Minimal())
	if strings.TrimSpace(s.View()) == "" {
		t.Error("expected non-empty now view")
	}
}

func TestContact_RendersContent(t *testing.T) {
	s := sections.NewContact(styles.Minimal())
	if strings.TrimSpace(s.View()) == "" {
		t.Error("expected non-empty contact view")
	}
}

func TestResume_RendersContent(t *testing.T) {
	s := sections.NewResume(styles.Minimal())
	if strings.TrimSpace(s.View()) == "" {
		t.Error("expected non-empty resume view")
	}
}
