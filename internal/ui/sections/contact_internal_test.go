package sections

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

func TestContact_DoneViewShowsName(t *testing.T) {
	c := NewContact(styles.Minimal())
	c.state = contactDone
	c.data.name = "Alice"
	c.data.email = "alice@test.com"
	if !strings.Contains(c.View().Content, "Alice") {
		t.Error("expected done view to contain submitter name")
	}
}

func TestContact_DoneViewShowsMessageSent(t *testing.T) {
	c := NewContact(styles.Minimal())
	c.state = contactDone
	c.data.name = "Alice"
	if !strings.Contains(c.View().Content, "message sent") {
		t.Error("expected done view to contain 'message sent'")
	}
}

func TestContact_DoneKeyReturnsToLinks(t *testing.T) {
	c := NewContact(styles.Minimal())
	c.state = contactDone
	c.data.name = "Alice"
	next, _ := c.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	if next.(Contact).state != contactLinks {
		t.Error("expected state to return to contactLinks after key press in done state")
	}
}

func TestContact_FormStateOnM(t *testing.T) {
	c := NewContact(styles.Minimal())
	c.width = 80
	next, _ := c.Update(tea.KeyPressMsg{Text: "m"})
	if next.(Contact).state != contactForm {
		t.Error("expected state to be contactForm after pressing m")
	}
}

func TestContact_AbortReturnsToLinks(t *testing.T) {
	c := NewContact(styles.Minimal())
	c.state = contactForm
	// Set huh form to aborted state directly — simulates the user pressing ESC
	c.form.State = huh.StateAborted
	// Any message triggers the post-update state check in our Update handler
	next, _ := c.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	if next.(Contact).state != contactLinks {
		t.Error("expected state to return to contactLinks on form abort")
	}
}

func TestContact_FormDataHasEmailField(t *testing.T) {
	c := NewContact(styles.Minimal())
	c.data.email = "test@example.com"
	if c.data.email != "test@example.com" {
		t.Error("expected formData to have an email field")
	}
}
