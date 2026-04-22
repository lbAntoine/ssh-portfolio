package sections

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

// Now displays what Antoine is currently working on.
type Now struct{ theme styles.Theme }

// NewNow returns an initialized Now section.
func NewNow(theme styles.Theme) Now { return Now{theme: theme} }

// Init implements tea.Model.
func (n Now) Init() tea.Cmd { return nil }

// Update implements tea.Model.
func (n Now) Update(_ tea.Msg) (tea.Model, tea.Cmd) { return n, nil }

// View implements tea.Model.
func (n Now) View() string {
	t := n.theme
	var b strings.Builder

	b.WriteString(t.Title.Render("now") + "\n\n")

	items := []struct{ icon, text string }{
		{"⚙  ", "Developer & test analyst at Nexpublica"},
		{"🌱 ", "Maintaining and improving Compendium"},
		{"🃏 ", "Validating TCG tournament platform with stores and TOs"},
		{"📖 ", "Finishing software architecture master's degree (Oct. 2026)"},
		{"🇰🇷 ", "Learning Korean — 안녕하세요"},
	}

	for _, item := range items {
		b.WriteString(t.Accent.Render(item.icon) + t.Body.Render(item.text) + "\n")
	}

	return t.Container.Render(b.String())
}
