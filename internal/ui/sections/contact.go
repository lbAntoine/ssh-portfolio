package sections

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

// Contact displays the about me section
type Contact struct{ theme styles.Theme }

// NewContact returns an initialized Contact section
func NewContact(theme styles.Theme) Contact { return Contact{theme: theme} }

// Init implements tea.Model
func (c Contact) Init() tea.Cmd { return nil }

// Update implements tea.Model
func (c Contact) Update(_ tea.Msg) (tea.Model, tea.Cmd) { return c, nil }

// View implements tea.Model
func (c Contact) View() string { return c.theme.Muted.Render("[ contact -- coming soon ]") }
