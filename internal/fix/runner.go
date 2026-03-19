package fix

import "fmt"

func Run(name string) error {
	fn, ok := Get(name)
	if !ok {
		return fmt.Errorf("unknown fix: %s", name)
	}

	return fn()
}
