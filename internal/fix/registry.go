package fix

type FixFunc func() error

var registry = make(map[string]FixFunc)

func Register(name string, fn FixFunc) {
	registry[name] = fn
}

func Get(name string) (FixFunc, bool) {
	fn, ok := registry[name]
	return fn, ok
}

func List() []string {
	var names []string
	for k := range registry {
		names = append(names, k)
	}
	return names
}
