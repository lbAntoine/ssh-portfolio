package sections

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

// About displays the about me section
type About struct{ theme styles.Theme }

// NewAbout returns an initialized About section
func NewAbout(theme styles.Theme) About { return About{theme: theme} }

// Init implements tea.Model
func (a About) Init() tea.Cmd { return nil }

// Update implements tea.Model
func (a About) Update(_ tea.Msg) (tea.Model, tea.Cmd) { return a, nil }

// View implements tea.Model
func (a About) View() string { return a.theme.Muted.Render("[ about -- coming soon ]") }
