package sections

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

type project struct {
	name   string
	status string
	desc   string
	url    string
}

var projectList = []project{
	{
		name:   "Compendium",
		status: "live",
		desc:   "Card lending & collection platform for the Flesh and Blood TCG community",
		url:    "compendium.vaultofsuraya.com",
	},
	{
		name:   "TCG Tournament Platform",
		status: "in development",
		desc:   "Tournament management, meta reporting and branded inforgraphic generation for TCG stores and TOs",
		url:    "",
	},
	{
		name:   "ssh portfolio",
		status: "you are here",
		desc:   "Interactive terminal portfolio served over SSH. Built with Go + Charmbracelet",
		url:    "ssh antoinelb.fr -p 2222",
	},
}

// Project displays the about me section
type Project struct {
	theme    styles.Theme
	selected int
	width    int
	height   int
}

// SetSize updates the dimensions of the section
func (p *Project) SetSize(width, height int) {
	p.width = width
	p.height = height
}

// NewProject returns an initialized Project section
func NewProject(theme styles.Theme) Project { return Project{theme: theme} }

// Init implements tea.Model
func (p Project) Init() tea.Cmd { return nil }

// Update implements tea.Model
func (p Project) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch string(msg.Runes) {
		case "j":
			if p.selected < len(projectList)-1 {
				p.selected++
			}
		case "k":
			if p.selected > 0 {
				p.selected--
			}
		}
	}
	return p, nil
}

// View implements tea.Model
func (p Project) View() string {
	t := p.theme
	var b strings.Builder

	b.WriteString(t.Title.Render("projects") + "\n\n")

	for i, proj := range projectList {
		cursor := "  "
		name := t.Body.Render(proj.name)
		status := t.Muted.Render("· " + proj.status)

		if i == p.selected {
			cursor = t.Accent.Render("▶ ")
			name = t.Accent.Render(proj.name)
			b.WriteString(cursor + name + " " + status + "\n")
			b.WriteString("    " + t.Body.Render(proj.desc) + "\n")
			if proj.url != "" {
				b.WriteString("    " + t.Muted.Render(proj.url) + "\n")
			}
			b.WriteString("\n")
		} else {
			b.WriteString(cursor + name + " " + status + "\n\n")
		}
	}

	b.WriteString(t.Muted.Render("j/k to navigate"))
	return t.Container.Width(p.width).Render(b.String())
}
