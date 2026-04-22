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

func TestAbout_RendersLocation(t *testing.T) {
	s := sections.NewAbout(styles.Minimal())
	if !strings.Contains(s.View(), "Aix-en-Provence") {
		t.Error("expected location in about view")
	}
}

func TestAbout_RendersRole(t *testing.T) {
	s := sections.NewAbout(styles.Minimal())
	if !strings.Contains(s.View(), "Backend") {
		t.Error("expected role in about view")
	}
}

func TestProjects_RendersContent(t *testing.T) {
	s := sections.NewProject(styles.Minimal())
	if strings.TrimSpace(s.View()) == "" {
		t.Error("expected non-empty projects view")
	}
}

func TestProjects_RendersCompendium(t *testing.T) {
	s := sections.NewProject(styles.Minimal())
	if !strings.Contains(s.View(), "Compendium") {
		t.Error("expected Compendium in project view")
	}
}

func TestProjects_RendersTCGPlatform(t *testing.T) {
	s := sections.NewProject(styles.Minimal())
	if !strings.Contains(s.View(), "TCG") {
		t.Error("expected TCG in project view")
	}
}

func TestProjects_RendersSSHPortfolio(t *testing.T) {
	s := sections.NewProject(styles.Minimal())
	if !strings.Contains(s.View(), "ssh") {
		t.Error("expected SSH in project view")
	}
}

func TestStack_RendersContent(t *testing.T) {
	s := sections.NewStack(styles.Minimal())
	if strings.TrimSpace(s.View()) == "" {
		t.Error("expected non-empty stack view")
	}
}

func TestStack_RendersLanguages(t *testing.T) {
	s := sections.NewStack(styles.Minimal())
	for _, lang := range []string{"Go", "Typescript", "Java", "Elixir"} {
		if !strings.Contains(s.View(), lang) {
			t.Errorf("expected %q in stack view", lang)
		}
	}
}

func TestStack_RendersInfra(t *testing.T) {
	s := sections.NewStack(styles.Minimal())
	for _, tool := range []string{"Docker", "PostgreSQL", "MongoDB", "Linux"} {
		if !strings.Contains(s.View(), tool) {
			t.Errorf("expected %q in stack view", tool)
		}
	}
}

func TestNow_RendersCurrentWork(t *testing.T) {
	s := sections.NewNow(styles.Minimal())
	for _, item := range []string{"Nexpublica", "Compendium"} {
		if !strings.Contains(s.View(), item) {
			t.Errorf("expected %q in now view", item)
		}
	}
}

func TestNow_RendersLinks(t *testing.T) {
	s := sections.NewContact(styles.Minimal())
	for _, item := range []string{"github", "linkedin", "antoine.lebras"} {
		if !strings.Contains(s.View(), item) {
			t.Errorf("expected %q in now view", item)
		}
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

func TestResume_ShowURL(t *testing.T) {
	s := sections.NewResume(styles.Minimal())
	if !strings.Contains(s.View(), "github.com/lbAntoine/ssh-portfolio") {
		t.Error("expected resume URL in resume view")
	}
}

func TestResume_RendersContent(t *testing.T) {
	s := sections.NewResume(styles.Minimal())
	if strings.TrimSpace(s.View()) == "" {
		t.Error("expected non-empty resume view")
	}
}
