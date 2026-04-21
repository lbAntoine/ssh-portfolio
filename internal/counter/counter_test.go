package counter_test

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/lbAntoine/ssh-portfolio/internal/counter"
)

func tempPath(t *testing.T) string {
	t.Helper()
	return filepath.Join(t.TempDir(), "counter.json")
}

func TestCounter_FirstIncrementReturnsOne(t *testing.T) {
	c := counter.New(tempPath(t))
	n, err := c.Increment()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 1 {
		t.Errorf("expected 1, got %d", n)
	}
}

func TestCounter_IncrementPersistsAcrossInstances(t *testing.T) {
	path := tempPath(t)
	c1 := counter.New(path)
	if _, err := c1.Increment(); err != nil {
		t.Fatal(err)
	}

	c2 := counter.New(path)
	n, err := c2.Current()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 1 {
		t.Errorf("expected 1, got %d", n)
	}
}

func TestCounter_ConcurrentIncrements(t *testing.T) {
	c := counter.New(tempPath(t))
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := c.Increment(); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}()
	}
	wg.Wait()

	n, err := c.Current()
	if err != nil {
		t.Fatal(err)
	}
	if n != 100 {
		t.Errorf("expected 100, got %d", n)
	}
}

func TestCounter_CorruptFileRecovers(t *testing.T) {
	path := tempPath(t)
	if err := os.WriteFile(path, []byte("not json"), 0o644); err != nil {
		t.Fatal(err)
	}

	c := counter.New(path)
	n, err := c.Increment()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 1 {
		t.Errorf("expected 1 after recovery, got %d", n)
	}
}
