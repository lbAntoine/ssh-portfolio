package sections

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

type nowGroup struct {
	label string
	items []string
}

var nowGroups = []nowGroup{
	{
		label: "work",
		items: []string{
			"Developer & test analyst at Nexpublica",
			"Finishing software architecture master's degree (Oct. 2026)",
		},
	},
	{
		label: "building",
		items: []string{
			"Maintaining and improving Compendium",
			"Validating TCG tournament platform with stores and TOs",
		},
	},
	{
		label: "learning",
		items: []string{
			"Korean — 안녕하세요",
		},
	},
}

// Now displays what Antoine is currently working on.
type Now struct {
	theme  styles.Theme
	width  int
	height int
}

// SetSize updates the dimensions of the section
func (n *Now) SetSize(width, height int) {
	n.width = width
	n.height = height
}

// NewNow returns an initialized Now section.
func NewNow(theme styles.Theme) Now { return Now{theme: theme} }

// Init implements tea.Model.
func (n Now) Init() tea.Cmd { return nil }

// Update implements tea.Model.
func (n Now) Update(_ tea.Msg) (tea.Model, tea.Cmd) { return n, nil }

// View implements tea.Model.
func (n Now) View() tea.View {
	t := n.theme
	var b strings.Builder

	b.WriteString(t.Title.Render("now") + "\n\n")

	for _, group := range nowGroups {
		b.WriteString(t.Accent.Render(group.label) + "\n")
		for _, item := range group.items {
			b.WriteString(t.Body.Render("  "+item) + "\n")
		}
		b.WriteString("\n")
	}

	b.WriteString(t.Muted.Render("updated april 2026"))

	return tea.NewView(t.Container.Width(n.width).Render(b.String()))
}
