package sections

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

// Project displays the about me section
type Project struct{ theme styles.Theme }

// NewProject returns an initialized Project section
func NewProject(theme styles.Theme) Project { return Project{theme: theme} }

// Init implements tea.Model
func (p Project) Init() tea.Cmd { return nil }

// Update implements tea.Model
func (p Project) Update(_ tea.Msg) (tea.Model, tea.Cmd) { return p, nil }

// View implements tea.Model
func (p Project) View() string { return p.theme.Muted.Render("[ project -- coming soon ]") }
