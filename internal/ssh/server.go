package ssh

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/log/v2"
	cssh "github.com/charmbracelet/ssh"
	"charm.land/wish/v2"
	at "charm.land/wish/v2/activeterm"
	bm "charm.land/wish/v2/bubbletea"
	lm "charm.land/wish/v2/logging"

	"github.com/lbAntoine/ssh-portfolio/internal/counter"
	"github.com/lbAntoine/ssh-portfolio/internal/ui"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

func makeHandler(c *counter.Counter) bm.Handler {
	return func(s cssh.Session) (tea.Model, []tea.ProgramOption) {
		theme := styles.Minimal()
		n, err := c.Increment()
		if err != nil {
			log.Warn("could not increment counter", "err", err)
			n = 0
		}
		return ui.NewApp(theme, n), nil
	}
}

// NewServer creates and configure a Wish SSH server on the given address
// using the provided host key path
func NewServer(addr, hostKeyPath string, c *counter.Counter) *cssh.Server {
	s, err := wish.NewServer(
		wish.WithAddress(addr),
		wish.WithHostKeyPath(hostKeyPath),
		wish.WithMiddleware(
			bm.Middleware(makeHandler(c)),
			at.Middleware(),
			lm.Middleware(),
		),
		wish.WithPublicKeyAuth(func(_ cssh.Context, _ cssh.PublicKey) bool {
			return true
		}),
		wish.WithIdleTimeout(5*time.Minute),
		wish.WithMaxTimeout(30*time.Minute),
	)
	if err != nil {
		log.Error("could not create server", "err", err)
		return nil
	}
	return s
}
