package styles

import gloss "charm.land/lipgloss/v2"

// Minimal theme configuration
func Minimal() Theme {
	return Theme{
		Name:      "minimal",
		Title:     gloss.NewStyle().Bold(true).Foreground(gloss.Color("78")),
		Subtitle:  gloss.NewStyle().Foreground(gloss.Color("250")),
		Body:      gloss.NewStyle().Foreground(gloss.Color("245")),
		Accent:    gloss.NewStyle().Bold(true).Foreground(gloss.Color("78")),
		Muted:     gloss.NewStyle().Foreground(gloss.Color("72")),
		Container: gloss.NewStyle().Padding(1, 2),
		Divider:   "────────────────────────────",
	}
}
