package sections

import (
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
	"github.com/lbAntoine/ssh-portfolio/internal/notifier"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

type contactState int

const (
	contactLinks contactState = iota
	contactForm
	contactDone
)

type formData struct {
	name    string
	email   string
	message string
}

// link is the type for links
type link struct {
	label string
	value string
}

// Contact displays contact links and an optional message form.
type Contact struct {
	theme  styles.Theme
	width  int
	height int
	state  contactState
	form   *huh.Form
	data   *formData
}

func newContactForm(data *formData) *huh.Form {
	f := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("name").
				Placeholder("your name").
				Value(&data.name),
			huh.NewInput().
				Title("email").
				Placeholder("your email").
				Value(&data.email),
			huh.NewText().
				Title("message").
				Description("alt+enter / ctrl+j for new line").
				Placeholder("what's on your mind?").
				Lines(4).
				CharLimit(500).
				Value(&data.message),
		),
	)
	f.SubmitCmd = nil
	f.CancelCmd = nil
	return f
}

// NewContact returns an initialized Contact section.
func NewContact(theme styles.Theme) Contact {
	data := &formData{}
	return Contact{theme: theme, form: newContactForm(data), data: data}
}

// SetSize updates the dimensions of the section.
func (c *Contact) SetSize(width, height int) {
	c.width = width
	c.height = height
}

// InputConsuming returns true when the huh form is active, signaling the
// parent model to bypass global key bindings and forward input directly here.
func (c Contact) InputConsuming() bool {
	return c.state == contactForm
}

// Init implements tea.Model.
func (c Contact) Init() tea.Cmd { return nil }

// Update implements tea.Model.
func (c Contact) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if wm, ok := msg.(tea.WindowSizeMsg); ok {
		cw := wm.Width - 4
		if cw > 96 {
			cw = 96
		}
		c.width = cw
		return c, nil
	}

	switch c.state {
	case contactLinks:
		if kp, ok := msg.(tea.KeyPressMsg); ok && kp.String() == "m" {
			w := c.width
			if w <= 0 {
				w = 80
			}
			c.form = c.form.WithWidth(w)
			c.state = contactForm
			return c, c.form.Init()
		}
	case contactForm:
		next, cmd := c.form.Update(msg)
		if f, ok := next.(*huh.Form); ok {
			c.form = f
		}
		switch c.form.State {
		case huh.StateCompleted:
			if url := os.Getenv("DISCORD_WEBHOOK_URL"); url != "" {
				notifier.SendDiscord(url, c.data.name, c.data.email, c.data.message)
			}
			c.state = contactDone
		case huh.StateAborted:
			c.data = &formData{}
			c.form = newContactForm(c.data)
			c.state = contactLinks
		}
		return c, cmd
	case contactDone:
		if _, ok := msg.(tea.KeyPressMsg); ok {
			c.data = &formData{}
			c.form = newContactForm(c.data)
			c.state = contactLinks
			return c, nil
		}
	}
	return c, nil
}

// View implements tea.Model.
func (c Contact) View() tea.View {
	t := c.theme
	switch c.state {
	case contactForm:
		return tea.NewView(c.form.View())
	case contactDone:
		var b strings.Builder
		b.WriteString(t.Title.Render("message sent") + "\n\n")
		b.WriteString(t.Body.Render("thanks, "+c.data.name+".") + "\n")
		b.WriteString(t.Muted.Render("i'll get back to you soon.") + "\n\n")
		b.WriteString(t.Muted.Render("press any key to go back"))
		return tea.NewView(t.Container.Width(c.width).Render(b.String()))
	default:
		var b strings.Builder
		b.WriteString(t.Title.Render("contact") + "\n\n")
		for _, l := range []link{
			{"github  ", "github.com/lbAntoine"},
			{"linkedin", "linkedin.com/in/antoine-le-bras"},
			{"email   ", "antoine.lebras+sshportfolio@gmail.com"},
			{"ssh     ", "ssh antoinelb.fr -p 2222"},
		} {
			b.WriteString(t.Accent.Render(l.label) + "  " + t.Muted.Render(l.value) + "\n")
		}
		b.WriteString("\n" + t.Muted.Render("press m to send a message"))
		return tea.NewView(t.Container.Width(c.width).Render(b.String()))
	}
}
