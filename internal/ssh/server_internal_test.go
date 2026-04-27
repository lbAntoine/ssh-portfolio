package ssh

import (
	"testing"

	"github.com/lbAntoine/ssh-portfolio/internal/counter"
)

func TestMakeHandler_Success(t *testing.T) {
	c := counter.New(t.TempDir() + "/c.json")
	h := makeHandler(c)
	model, opts := h(nil)
	if model == nil {
		t.Error("expected non-nil model")
	}
	if opts != nil {
		t.Error("expected nil program options")
	}
}

func TestMakeHandler_CounterError(t *testing.T) {
	// Path in a non-existent subdirectory causes Increment to fail at write time
	c := counter.New(t.TempDir() + "/missing/c.json")
	h := makeHandler(c)
	model, _ := h(nil)
	if model == nil {
		t.Error("expected non-nil model even when counter fails")
	}
}
