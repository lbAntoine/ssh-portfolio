package sections

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
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
type Stack struct{ theme styles.Theme }

// NewStack returns an initialized Stack section.
func NewStack(theme styles.Theme) Stack { return Stack{theme: theme} }

// Init implements tea.Model.
func (s Stack) Init() tea.Cmd { return nil }

// Update implements tea.Model.
func (s Stack) Update(_ tea.Msg) (tea.Model, tea.Cmd) { return s, nil }

// View implements tea.Model.
func (s Stack) View() string {
	t := s.theme
	var b strings.Builder

	b.WriteString(t.Title.Render("stack") + "\n\n")

	for _, cat := range stackCategories {
		b.WriteString(t.Accent.Render(cat.label) + "\n")
		b.WriteString(t.Muted.Render("  "+strings.Join(cat.items, " · ")) + "\n\n")
	}

	return t.Container.Render(b.String())
}
