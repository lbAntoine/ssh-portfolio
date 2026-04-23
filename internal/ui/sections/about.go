package sections

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

// About displays the about me section
type About struct {
	theme  styles.Theme
	width  int
	height int
}

// SetSize updates the dimensions of the section
func (a *About) SetSize(width, height int) {
	a.width = width
	a.height = height
}

// NewAbout returns an initialized About section
func NewAbout(theme styles.Theme) About { return About{theme: theme} }

// Init implements tea.Model
func (a About) Init() tea.Cmd { return nil }

// Update implements tea.Model
func (a About) Update(_ tea.Msg) (tea.Model, tea.Cmd) { return a, nil }

// View implements tea.Model
func (a About) View() string {
	t := a.theme
	return t.Container.Width(a.width).Render(
		t.Title.Render("about") + "\n\n" +
			t.Body.Render("Backend developer finishing a software architecture master's degree (Oct. 2026).") + "\n" +
			t.Body.Render("I build tools I want to use — for communities I'm part of.") + "\n\n" +
			t.Muted.Render(t.Divider) + "\n\n" +
			t.Accent.Render("⚙  ") + t.Body.Render("Go · Typescript · Java — learning Elixir") + "\n" +
			t.Accent.Render("📍 ") + t.Body.Render("Aix-en-Provence, France") + "\n" +
			t.Accent.Render("❤  ") + t.Body.Render("Open source, TCG scene, building things that solve real problems") + "\n" +
			t.Accent.Render("⚡ ") + t.Body.Render("Card game enthusiast, keyboard hobbyist, gunpla builder"),
	)
}
