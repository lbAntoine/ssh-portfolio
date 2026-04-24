package sections

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/viewport"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

const resumeURL = "https://github.com/lbAntoine/ssh-portfolio/releases/latest/download/resume.pdf"

// Resume displays the resume download section.
type Resume struct {
	theme  styles.Theme
	vp     viewport.Model
	width  int
	height int
}

// SetSize updates the dimensions of the section
func (r *Resume) SetSize(width, height int) {
	r.width = width
	r.height = height
	r.vp.SetWidth(width)
	r.vp.SetHeight(height)
	r.vp.SetContent(r.renderContent())
}

func (r Resume) renderContent() string {
	t := r.theme
	var b strings.Builder
	b.WriteString(t.Title.Render("resume") + "\n\n")
	b.WriteString(t.Body.Render("Download my resume at:") + "\n\n")
	b.WriteString(t.Accent.Render(resumeURL) + "\n\n")
	b.WriteString(t.Muted.Render("tip: copy the URL above and paste it in your browser"))
	return t.Container.Width(r.width).Render(b.String())
}

// NewResume returns an initialized Resume section.
func NewResume(theme styles.Theme) Resume {
	vp := viewport.New()
	return Resume{theme: theme, vp: vp}
}

// Init implements tea.Model.
func (r Resume) Init() tea.Cmd { return nil }

// Update implements tea.Model.
func (r Resume) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	r.vp, cmd = r.vp.Update(msg)
	return r, cmd
}

// View implements tea.Model.
func (r Resume) View() tea.View {
	if r.width == 0 {
		return tea.NewView(r.renderContent())
	}
	return tea.NewView(r.vp.View())
}
