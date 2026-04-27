package ssh_test

import (
	"testing"

	"github.com/lbAntoine/ssh-portfolio/internal/counter"
	"github.com/lbAntoine/ssh-portfolio/internal/ssh"
)

func TestServer_Configuration(t *testing.T) {
	c := counter.New(t.TempDir() + "/counter.json")
	srv := ssh.NewServer(":2222", "./data/host_key", c)
	if srv == nil {
		t.Fatal("expected a non-nil server")
	}
}

func TestServer_InvalidKeyPath(t *testing.T) {
	c := counter.New(t.TempDir() + "/counter.json")
	srv := ssh.NewServer(":2222", "/nonexistent_dir/host_key", c)
	if srv != nil {
		t.Error("expected nil server when key path is invalid")
	}
}
