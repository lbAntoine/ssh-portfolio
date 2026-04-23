package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"
	"github.com/lbAntoine/ssh-portfolio/internal/counter"
	sshsrv "github.com/lbAntoine/ssh-portfolio/internal/ssh"
	"github.com/lbAntoine/ssh-portfolio/internal/ui/styles"
)

func main() {
	port := flag.String("port", "2222", "SSH server port")
	hostKey := flag.String("host-key", "./data/host_key", "path to host key")
	counterPath := flag.String("counter", "./data/counter.json", "path to visitor counter file")
	logLevel := flag.String("log-level", "info", "log level (debug, info, warn, error)")
	flag.Parse()

	if lvl, err := log.ParseLevel(*logLevel); err == nil {
		log.SetLevel(lvl)
	}

	addr := ":" + *port
	c := counter.New(*counterPath)

	srv := sshsrv.NewServer(addr, *hostKey, styles.Minimal(), c)
	if srv == nil {
		log.Error("failed to create server")
		os.Exit(1)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	log.Info("starting ssh-portfolio", "addr", addr, "host-key", *hostKey, "counter", *counterPath)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("server error", "err", err)
			done <- syscall.SIGTERM
		}
	}()

	sig := <-done
	log.Info("shutting down", "signal", sig)
}
