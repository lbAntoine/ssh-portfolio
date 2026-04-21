package ssh_test

import (
	"testing"

	"github.com/lbAntoine/ssh-portfolio/internal/ssh"
)

func TestServer_Configuration(t *testing.T) {
	srv := ssh.NewServer(":2222", "./data/host_key")
	if srv == nil {
		t.Fatal("expected a non-nil server")
	}
}
