package checks

import "strings"
import "github.com/paszed/doctor/internal/model"

type CheckFunc func() model.Result

var registry = map[string]CheckFunc{}

func Register(name string, fn CheckFunc) {
	registry[strings.ToLower(name)] = fn
}
