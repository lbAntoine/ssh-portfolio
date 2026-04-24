package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/splash"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

// App is the top level model managing splash -> main transition (and others in the future)
type App struct {
	splash splash.Model
	main   Model
	ready  bool
}

// NewApp returns an initialized App starting in splash state
func NewApp(theme styles.Theme, visitorCount int) App {
	return App{
		splash: splash.NewSplash(splash.SplashText, theme),
		main:   NewModel(theme, visitorCount),
	}
}

// Ready returns true when the splash is done and the main model is active
func (a App) Ready() bool { return a.ready }

// Init Implements tea.Model
func (a App) Init() tea.Cmd { return a.splash.Init() }

// Update implements tea.Model
func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if wm, ok := msg.(tea.WindowSizeMsg); ok {
		splashNext, _ := a.splash.Update(wm)
		a.splash = splashNext.(splash.Model)
		mainNext, _ := a.main.Update(wm)
		a.main = mainNext.(Model)
		return a, nil
	}

	if a.ready {
		next, cmd := a.main.Update(msg)
		a.main = next.(Model)
		return a, cmd
	}

	next, cmd := a.splash.Update(msg)
	a.splash = next.(splash.Model)
	if a.splash.Done() {
		a.ready = true
	}
	return a, cmd
}

// View implements tea.Model
func (a App) View() string {
	if a.ready {
		return a.main.View()
	}
	return a.splash.View()
}
