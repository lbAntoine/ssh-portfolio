package styles_test

import (
	"testing"

	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

func TestBreakpoint_CompactUnder60(t *testing.T) {
	bp := styles.BreakpointFor(59)
	if bp != styles.Compact {
		t.Errorf("expected Compact for width 59, got %v", bp)
	}
}

func TestBreakpoint_StandardAt80(t *testing.T) {
	bp := styles.BreakpointFor(80)
	if bp != styles.Standard {
		t.Errorf("expected Standard for width 80, got %v", bp)
	}
}

func TestBreakpoint_WideAt120(t *testing.T) {
	bp := styles.BreakpointFor(120)
	if bp != styles.Wide {
		t.Errorf("expected Wide for width 120, got %v", bp)
	}
}
