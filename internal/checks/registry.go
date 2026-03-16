package checks

import "github.com/paszed/doctor/internal/model"

type CheckFunc func() model.Result

var Registry = []CheckFunc{
	GitCheck,
	NodeCheck,
	PythonCheck,
	DockerCheck,
}
