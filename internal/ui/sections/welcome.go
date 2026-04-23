package sections

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

// Welcome is the landing section of the portfolio
type Welcome struct {
	theme        styles.Theme
	visitorCount int
	width        int
	height       int
}

// SetSize updates the dimensions of the section
func (w *Welcome) SetSize(width, height int) {
	w.width = width
	w.height = height
}

// NewWelcome returns an initialized Welcome section
func NewWelcome(theme styles.Theme, visitorCount int) Welcome {
	return Welcome{theme: theme, visitorCount: visitorCount}
}

// Init implements tea.Model
func (w Welcome) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model
func (w Welcome) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return w, nil
}

// View implements tea.Model
func (w Welcome) View() string {
	t := w.theme

	name := t.Title.Render("Antoine Le Bras")
	subtitle := t.Subtitle.Render("Backend Developer · Aix-en-Provence, France")
	visitor := t.Muted.Render(fmt.Sprintf("you are visitor #%d", w.visitorCount))

	var content string
	if styles.BreakpointFor(w.width) == styles.Compact {
		content = name + "\n" + subtitle + "\n\n" + visitor
	} else {
		tagline := t.Body.Render("I build things I would use myself.")
		greeting := t.Muted.Render("안녕하세요")
		divider := t.Muted.Render(t.Divider)
		content = name + "\n" + subtitle + "\n\n" + tagline + "\n\n" + divider + "\n\n" + greeting + "  " + visitor
	}

	return t.Container.Render(content)
}
