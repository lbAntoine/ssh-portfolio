package sections_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/sections"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

func TestAbout_RendersContent(t *testing.T) {
	s := sections.NewAbout(styles.Minimal())
	if strings.TrimSpace(s.View().Content) == "" {
		t.Error("expected non-empty about view")
	}
}

func TestAbout_RendersLocation(t *testing.T) {
	s := sections.NewAbout(styles.Minimal())
	if !strings.Contains(s.View().Content, "Aix-en-Provence") {
		t.Error("expected location in about view")
	}
}

func TestAbout_RendersRole(t *testing.T) {
	s := sections.NewAbout(styles.Minimal())
	if !strings.Contains(s.View().Content, "Backend") {
		t.Error("expected role in about view")
	}
}

func TestProjects_RendersContent(t *testing.T) {
	s := sections.NewProject(styles.Minimal())
	if strings.TrimSpace(s.View().Content) == "" {
		t.Error("expected non-empty projects view")
	}
}

func TestProjects_RendersCompendium(t *testing.T) {
	s := sections.NewProject(styles.Minimal())
	if !strings.Contains(s.View().Content, "Compendium") {
		t.Error("expected Compendium in project view")
	}
}

func TestProjects_RendersTCGPlatform(t *testing.T) {
	s := sections.NewProject(styles.Minimal())
	if !strings.Contains(s.View().Content, "TCG") {
		t.Error("expected TCG in project view")
	}
}

func TestProjects_RendersSSHPortfolio(t *testing.T) {
	s := sections.NewProject(styles.Minimal())
	if !strings.Contains(s.View().Content, "ssh") {
		t.Error("expected SSH in project view")
	}
}

func TestStack_RendersContent(t *testing.T) {
	s := sections.NewStack(styles.Minimal())
	if strings.TrimSpace(s.View().Content) == "" {
		t.Error("expected non-empty stack view")
	}
}

func TestStack_RendersLanguages(t *testing.T) {
	s := sections.NewStack(styles.Minimal())
	for _, lang := range []string{"Go", "Typescript", "Java", "Elixir"} {
		if !strings.Contains(s.View().Content, lang) {
			t.Errorf("expected %q in stack view", lang)
		}
	}
}

func TestStack_RendersInfra(t *testing.T) {
	s := sections.NewStack(styles.Minimal())
	for _, tool := range []string{"Docker", "PostgreSQL", "MongoDB", "Linux"} {
		if !strings.Contains(s.View().Content, tool) {
			t.Errorf("expected %q in stack view", tool)
		}
	}
}

func TestNow_RendersCurrentWork(t *testing.T) {
	s := sections.NewNow(styles.Minimal())
	for _, item := range []string{"Nexpublica", "Compendium"} {
		if !strings.Contains(s.View().Content, item) {
			t.Errorf("expected %q in now view", item)
		}
	}
}

func TestNow_RendersLinks(t *testing.T) {
	s := sections.NewContact(styles.Minimal())
	for _, item := range []string{"github", "linkedin", "antoine.lebras"} {
		if !strings.Contains(s.View().Content, item) {
			t.Errorf("expected %q in now view", item)
		}
	}
}

func TestNow_RendersContent(t *testing.T) {
	s := sections.NewNow(styles.Minimal())
	if strings.TrimSpace(s.View().Content) == "" {
		t.Error("expected non-empty now view")
	}
}

func TestContact_RendersContent(t *testing.T) {
	s := sections.NewContact(styles.Minimal())
	if strings.TrimSpace(s.View().Content) == "" {
		t.Error("expected non-empty contact view")
	}
}

func TestResume_ShowURL(t *testing.T) {
	s := sections.NewResume(styles.Minimal())
	if !strings.Contains(s.View().Content, "github.com/lbAntoine/ssh-portfolio") {
		t.Error("expected resume URL in resume view")
	}
}

func TestResume_RendersContent(t *testing.T) {
	s := sections.NewResume(styles.Minimal())
	if strings.TrimSpace(s.View().Content) == "" {
		t.Error("expected non-empty resume view")
	}
}

// Init coverage — all sections return nil

func TestAbout_Init(t *testing.T) {
	if sections.NewAbout(styles.Minimal()).Init() != nil {
		t.Error("expected nil Init cmd")
	}
}

func TestContact_Init(t *testing.T) {
	if sections.NewContact(styles.Minimal()).Init() != nil {
		t.Error("expected nil Init cmd")
	}
}

func TestNow_Init(t *testing.T) {
	if sections.NewNow(styles.Minimal()).Init() != nil {
		t.Error("expected nil Init cmd")
	}
}

func TestProject_Init(t *testing.T) {
	if sections.NewProject(styles.Minimal()).Init() != nil {
		t.Error("expected nil Init cmd")
	}
}

func TestResume_Init(t *testing.T) {
	if sections.NewResume(styles.Minimal()).Init() != nil {
		t.Error("expected nil Init cmd")
	}
}

func TestStack_Init(t *testing.T) {
	if sections.NewStack(styles.Minimal()).Init() != nil {
		t.Error("expected nil Init cmd")
	}
}

func TestWelcome_Init(t *testing.T) {
	if sections.NewWelcome(styles.Minimal(), 0).Init() != nil {
		t.Error("expected nil Init cmd")
	}
}

// Update coverage — exercise the delegate path for each section

func TestAbout_Update(t *testing.T) {
	s := sections.NewAbout(styles.Minimal())
	s.SetSize(80, 20)
	_, _ = s.Update(tea.WindowSizeMsg{Width: 80, Height: 20})
}

func TestContact_Update(t *testing.T) {
	s := sections.NewContact(styles.Minimal())
	_, _ = s.Update(tea.WindowSizeMsg{})
}

func TestNow_Update(t *testing.T) {
	s := sections.NewNow(styles.Minimal())
	_, _ = s.Update(tea.WindowSizeMsg{})
}

func TestProject_Update(t *testing.T) {
	s := sections.NewProject(styles.Minimal())
	s.SetSize(80, 20)
	_, _ = s.Update(tea.KeyPressMsg{Code: tea.KeyDown})
}

func TestResume_Update(t *testing.T) {
	s := sections.NewResume(styles.Minimal())
	s.SetSize(80, 20)
	_, _ = s.Update(tea.WindowSizeMsg{Width: 80, Height: 20})
}

func TestStack_Update(t *testing.T) {
	s := sections.NewStack(styles.Minimal())
	_, _ = s.Update(tea.WindowSizeMsg{})
}

func TestWelcome_Update(t *testing.T) {
	s := sections.NewWelcome(styles.Minimal(), 0)
	_, _ = s.Update(tea.WindowSizeMsg{})
}

// SetSize coverage

func TestAbout_SetSize(t *testing.T) {
	s := sections.NewAbout(styles.Minimal())
	s.SetSize(80, 20)
}

func TestContact_SetSize(t *testing.T) {
	s := sections.NewContact(styles.Minimal())
	s.SetSize(80, 20)
}

func TestNow_SetSize(t *testing.T) {
	s := sections.NewNow(styles.Minimal())
	s.SetSize(80, 20)
}

func TestProject_SetSize(t *testing.T) {
	s := sections.NewProject(styles.Minimal())
	s.SetSize(80, 20)
}

func TestResume_SetSize(t *testing.T) {
	s := sections.NewResume(styles.Minimal())
	s.SetSize(80, 20)
}

func TestStack_SetSize(t *testing.T) {
	s := sections.NewStack(styles.Minimal())
	s.SetSize(80, 20)
}

// Branch coverage — viewport and compact paths

func TestAbout_ViewWithViewport(t *testing.T) {
	s := sections.NewAbout(styles.Minimal())
	// height=1 ensures content overflows → AtBottom()=false → scroll hint shown
	s.SetSize(80, 1)
	v := s.View()
	if strings.TrimSpace(v.Content) == "" {
		t.Error("expected non-empty about view after SetSize")
	}
}

func TestResume_ViewWithViewport(t *testing.T) {
	s := sections.NewResume(styles.Minimal())
	s.SetSize(80, 20)
	v := s.View()
	if strings.TrimSpace(v.Content) == "" {
		t.Error("expected non-empty resume view after SetSize")
	}
}

func TestContact_InputConsuming_DefaultFalse(t *testing.T) {
	s := sections.NewContact(styles.Minimal())
	if s.InputConsuming() {
		t.Error("expected InputConsuming to be false before entering form")
	}
}

func TestContact_InputConsuming_TrueAfterM(t *testing.T) {
	s := sections.NewContact(styles.Minimal())
	s.SetSize(80, 20)
	next, _ := s.Update(tea.KeyPressMsg{Text: "m"})
	if !next.(interface{ InputConsuming() bool }).InputConsuming() {
		t.Error("expected InputConsuming to be true after pressing m")
	}
}

func TestStack_CompactView(t *testing.T) {
	s := sections.NewStack(styles.Minimal())
	// width < 60 → Compact breakpoint → bullet-per-line layout
	s.SetSize(40, 20)
	v := s.View()
	if strings.TrimSpace(v.Content) == "" {
		t.Error("expected non-empty compact stack view")
	}
}
