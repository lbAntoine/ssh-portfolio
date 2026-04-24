package sections

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/list"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

type projectItem struct {
	name   string
	status string
	desc   string
	url    string
}

func (p projectItem) Title() string       { return p.name + "  · " + p.status }
func (p projectItem) Description() string {
	if p.url != "" {
		return p.desc + "\n" + p.url
	}
	return p.desc
}
func (p projectItem) FilterValue() string { return p.name }

var projectItems = []list.Item{
	projectItem{
		name:   "Compendium",
		status: "live",
		desc:   "Card lending & collection platform for the Flesh and Blood TCG community",
		url:    "compendium.vaultofsuraya.com",
	},
	projectItem{
		name:   "TCG Tournament Platform",
		status: "in development",
		desc:   "Tournament management, meta reporting and branded infographic generation for TCG stores and TOs",
	},
	projectItem{
		name:   "ssh portfolio",
		status: "you are here",
		desc:   "Interactive terminal portfolio served over SSH. Built with Go + Charmbracelet",
		url:    "ssh antoinelb.fr -p 2222",
	},
}

// Project displays the projects section
type Project struct {
	theme  styles.Theme
	list   list.Model
	width  int
	height int
}

// SetSize updates the dimensions of the section
func (p *Project) SetSize(width, height int) {
	p.width = width
	p.height = height
	p.list.SetWidth(width)
	p.list.SetHeight(height)
}

// NewProject returns an initialized Project section
func NewProject(theme styles.Theme) Project {
	l := list.New(projectItems, list.NewDefaultDelegate(), 80, 20)
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.SetShowPagination(false)
	l.SetShowHelp(false)
	l.SetFilteringEnabled(false)
	return Project{theme: theme, list: l}
}

// Init implements tea.Model
func (p Project) Init() tea.Cmd { return nil }

// Update implements tea.Model
func (p Project) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	p.list, cmd = p.list.Update(msg)
	return p, cmd
}

// View implements tea.Model
func (p Project) View() tea.View {
	t := p.theme
	header := t.Container.Width(p.width).Render(t.Title.Render("projects"))
	return tea.NewView(header + "\n" + p.list.View())
}
