package sections

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

// link is the type for links
type link struct {
	label string
	value string
}

// Contact displays contact links.
type Contact struct {
	theme  styles.Theme
	width  int
	height int
}

// SetSize updates the dimensions of the section
func (c *Contact) SetSize(width, height int) {
	c.width = width
	c.height = height
}

// NewContact returns an initialized Contact section.
func NewContact(theme styles.Theme) Contact { return Contact{theme: theme} }

// Init implements tea.Model.
func (c Contact) Init() tea.Cmd { return nil }

// Update implements tea.Model.
func (c Contact) Update(_ tea.Msg) (tea.Model, tea.Cmd) { return c, nil }

// View implements tea.Model.
func (c Contact) View() string {
	t := c.theme
	var b strings.Builder

	b.WriteString(t.Title.Render("contact") + "\n\n")

	links := []link{
		{"github  ", "github.com/lbAntoine"},
		{"linkedin", "linkedin.com/in/antoine-le-bras"},
		{"email   ", "antoine.lebras+sshportfolio@gmail.com"},
		{"here    ", "ssh antoinelb.fr -p 2222"},
	}

	for _, l := range links {
		b.WriteString(t.Accent.Render(l.label) + "  " + t.Body.Render(l.value) + "\n")
	}

	return t.Container.Render(b.String())
}
