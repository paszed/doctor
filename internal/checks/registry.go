package checks

import (
	"sort"

	"github.com/paszed/doctor/internal/model"
)

// CheckFunc defines a check function signature
type CheckFunc func() model.Result

// registry holds all registered checks by name
var registry = make(map[string]CheckFunc)

// Register adds a check to the registry
func Register(name string, fn CheckFunc) {
	registry[name] = fn
}

// Get returns a check by name
func Get(name string) (CheckFunc, bool) {
	fn, ok := registry[name]
	return fn, ok
}

// List returns all available check names (sorted)
func List() []string {
	var names []string

	for name := range registry {
		names = append(names, name)
	}

	sort.Strings(names)
	return names
}
