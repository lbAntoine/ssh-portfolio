package sections

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

type category struct {
	label string
	items []string
}

var stackCategories = []category{
	{label: "languages", items: []string{"Go", "Typescript", "Java", "Elixir (learning)"}},
	{label: "infra", items: []string{"Docker", "PostgreSQL", "MongoDB", "Coolify", "Linux"}},
	{label: "frontend", items: []string{"React", "TailwindCSS", "Vite"}},
}

// Stack displays the tech stack section.
type Stack struct {
	theme  styles.Theme
	width  int
	height int
}

// SetSize updates the dimensions of the section
func (s *Stack) SetSize(width, height int) {
	s.width = width
	s.height = height
}

// NewStack returns an initialized Stack section.
func NewStack(theme styles.Theme) Stack { return Stack{theme: theme} }

// Init implements tea.Model.
func (s Stack) Init() tea.Cmd { return nil }

// Update implements tea.Model.
func (s Stack) Update(_ tea.Msg) (tea.Model, tea.Cmd) { return s, nil }

// View implements tea.Model.
func (s Stack) View() tea.View {
	t := s.theme
	var b strings.Builder

	b.WriteString(t.Title.Render("stack") + "\n\n")

	compact := styles.BreakpointFor(s.width) == styles.Compact
	for _, cat := range stackCategories {
		b.WriteString(t.Accent.Render(cat.label) + "\n")
		if compact {
			for _, item := range cat.items {
				b.WriteString(t.Muted.Render("  · "+item) + "\n")
			}
			b.WriteString("\n")
		} else {
			b.WriteString(t.Muted.Render("  "+strings.Join(cat.items, " · ")) + "\n\n")
		}
	}

	return tea.NewView(t.Container.Width(s.width).Render(b.String()))
}
