package styles

import gloss "github.com/charmbracelet/lipgloss"

// Theme data type
type Theme struct {
	Name string

	// Text
	Title    gloss.Style
	Subtitle gloss.Style
	Body     gloss.Style
	Accent   gloss.Style
	Muted    gloss.Style

	// Layout
	Container gloss.Style
	Divider   string
}
