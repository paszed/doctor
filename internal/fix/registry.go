package fix

type FixFunc func(args []string) error

var registry = make(map[string]FixFunc)

// Register adds a fix function to the registry
func Register(name string, fn FixFunc) {
	registry[name] = fn
}

// Get returns a fix function by name
func Get(name string) (FixFunc, bool) {
	fn, ok := registry[name]
	return fn, ok
}

// List returns all registered fix names
func List() []string {
	var names []string
	for name := range registry {
		names = append(names, name)
	}
	return names
}
