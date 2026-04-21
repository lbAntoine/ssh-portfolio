package ssh_test

import (
	"testing"

	"github.com/lbAntoine/ssh-portfolio/internal/counter"
	"github.com/lbAntoine/ssh-portfolio/internal/ssh"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

func TestServer_Configuration(t *testing.T) {
	c := counter.New(t.TempDir() + "/counter.json")
	srv := ssh.NewServer(":2222", "./data/host_key", styles.Minimal(), c)
	if srv == nil {
		t.Fatal("expected a non-nil server")
	}
}
