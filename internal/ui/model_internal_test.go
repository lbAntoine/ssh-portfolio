package ui

import "testing"

func TestKeyMap_ShortHelp(t *testing.T) {
	if len(keys.ShortHelp()) == 0 {
		t.Error("expected non-empty short help bindings")
	}
}

func TestKeyMap_FullHelp(t *testing.T) {
	if len(keys.FullHelp()) == 0 {
		t.Error("expected non-empty full help rows")
	}
}
