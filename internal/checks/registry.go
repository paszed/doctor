package checks

import "github.com/paszed/doctor/internal/model"

type CheckFunc func() model.Result

var registry []CheckFunc

func Register(check CheckFunc) {
	registry = append(registry, check)
}

func All() []CheckFunc {
	return registry
}

func Names() []string {

	var names []string

	for _, check := range registry {

		result := check()
		names = append(names, result.Name)

	}

	return names
}
