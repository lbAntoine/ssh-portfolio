package ssh

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	cssh "github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"

	"github.com/lbAntoine/ssh-portfolio/internal/counter"
	"github.com/lbAntoine/ssh-portfolio/internal/ui"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

// NewServer creates and configure a Wish SSH server on the given address
// using the provided host key path
func NewServer(addr, hostKeyPath string, theme styles.Theme, c *counter.Counter) *cssh.Server {
	s, err := wish.NewServer(
		wish.WithAddress(addr),
		wish.WithHostKeyPath(hostKeyPath),
		wish.WithMiddleware(
			bm.Middleware(func(s cssh.Session) (tea.Model, []tea.ProgramOption) {
				n, err := c.Increment()
				if err != nil {
					log.Warn("could not increment counter", "err", err)
					n = 0
				}
				return ui.NewModel(theme, n), nil
			}),
			lm.Middleware(),
		),
		wish.WithPublicKeyAuth(func(_ cssh.Context, _ cssh.PublicKey) bool {
			return true
		}),
	)
	if err != nil {
		log.Error("could not create server", "err", err)
		return nil
	}
	return s
}
