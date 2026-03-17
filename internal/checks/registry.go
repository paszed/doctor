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
