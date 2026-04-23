package styles

// Breakpoint represents a terminal width category
type Breakpoint int

const (
	// Compact is used for terminals narrower than 60 col
	Compact Breakpoint = iota
	// Standard is used for terminals between 60 to 99 cols
	Standard
	// Wide is used for terminals larger than 120 col
	Wide
)

// BreakpointFor returns the Breakpoint for the given terminal width
func BreakpointFor(width int) Breakpoint {
	switch {
	case width < 60:
		return Compact
	case width < 100:
		return Standard
	default:
		return Wide
	}
}
