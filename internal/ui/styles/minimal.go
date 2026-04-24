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

// Minimal theme configuration
func MinimalWith(r *gloss.Renderer) Theme {
	return Theme{
		Name:      "minimal",
		Title:     r.NewStyle().Bold(true).Foreground(gloss.Color("78")),
		Subtitle:  r.NewStyle().Foreground(gloss.Color("250")),
		Body:      r.NewStyle().Foreground(gloss.Color("245")),
		Accent:    r.NewStyle().Bold(true).Foreground(gloss.Color("78")),
		Muted:     r.NewStyle().Foreground(gloss.Color("72")),
		Container: r.NewStyle().Padding(1, 2),
		Divider:   "────────────────────────────",
	}
}
