package sections

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

// Stack displays the about me section
type Stack struct{ theme styles.Theme }

// NewStack returns an initialized Stack section
func NewStack(theme styles.Theme) Stack { return Stack{theme: theme} }

// Init implements tea.Model
func (s Stack) Init() tea.Cmd { return nil }

// Update implements tea.Model
func (s Stack) Update(_ tea.Msg) (tea.Model, tea.Cmd) { return s, nil }

// View implements tea.Model
func (s Stack) View() string { return s.theme.Muted.Render("[ stack -- coming soon ]") }
