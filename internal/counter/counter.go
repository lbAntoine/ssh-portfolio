package counter

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/charmbracelet/log"
)

type data struct {
	Count int `json:"count"`
}

// Counter is a file-backed visitor counter safe for concurrent use
type Counter struct {
	mu   sync.Mutex
	path string
}

// New returns a Counter with a filepath
func New(path string) *Counter {
	return &Counter{path: path}
}

// Increment adds one to the counter and returns the new value
func (c *Counter) Increment() (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	d := c.load()
	d.Count++
	if err := c.save(d); err != nil {
		return 0, err
	}
	return d.Count, nil
}

// Current returns the current counter value without mutating it
func (c *Counter) Current() (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.load().Count, nil
}

// load reads the counter file, returning zero value on any error
func (c *Counter) load() data {
	b, err := os.ReadFile(c.path)
	if err != nil {
		return data{}
	}
	var d data
	if err := json.Unmarshal(b, &d); err != nil {
		log.Warn("counter file corrupt, resetting", "path", c.path)
		return data{}
	}
	return d
}

// save writes the counter atomically via a temp file + rename
func (c *Counter) save(d data) error {
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}
	tmp := c.path + ".tmp"
	if err := os.WriteFile(tmp, b, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, c.path)
}
