package sections

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

const resumeURL = "https://github.com/lbAntoine/ssh-portfolio/releases/latest/download/resume.pdf"

// Resume displays the resume download section.
type Resume struct{ theme styles.Theme }

// NewResume returns an initialized Resume section.
func NewResume(theme styles.Theme) Resume { return Resume{theme: theme} }

// Init implements tea.Model.
func (r Resume) Init() tea.Cmd { return nil }

// Update implements tea.Model.
func (r Resume) Update(_ tea.Msg) (tea.Model, tea.Cmd) { return r, nil }

// View implements tea.Model.
func (r Resume) View() string {
	t := r.theme
	var b strings.Builder

	b.WriteString(t.Title.Render("resume") + "\n\n")
	b.WriteString(t.Body.Render("Download my resume at:") + "\n\n")
	b.WriteString(t.Accent.Render(resumeURL) + "\n\n")
	b.WriteString(t.Muted.Render("tip: copy the URL above and paste it in your browser"))

	return t.Container.Render(b.String())
}
