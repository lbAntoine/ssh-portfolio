package sections

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

// Now displays the about me section
type Now struct{ theme styles.Theme }

// NewNow returns an initialized Now section
func NewNow(theme styles.Theme) Now { return Now{theme: theme} }

// Init implements tea.Model
func (n Now) Init() tea.Cmd { return nil }

// Update implements tea.Model
func (n Now) Update(_ tea.Msg) (tea.Model, tea.Cmd) { return n, nil }

// View implements tea.Model
func (n Now) View() string { return n.theme.Muted.Render("[ now -- coming soon ]") }
