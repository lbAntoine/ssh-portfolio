package sections

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

// Resume displays the about me section
type Resume struct{ theme styles.Theme }

// NewResume returns an initialized Resume section
func NewResume(theme styles.Theme) Resume { return Resume{theme: theme} }

// Init implements tea.Model
func (r Resume) Init() tea.Cmd { return nil }

// Update implements tea.Model
func (r Resume) Update(_ tea.Msg) (tea.Model, tea.Cmd) { return r, nil }

// View implements tea.Model
func (r Resume) View() string { return r.theme.Muted.Render("[ resume -- coming soon ]") }
