package checks

import (
	"context"
	"time"

	"github.com/paszed/doctor/internal/model"
)

// default timeout per check
const timeout = 2 * time.Second

// RunAll executes all checks concurrently with timeout protection
func RunAll() []model.Result {
	ch := make(chan model.Result)

	// launch all checks
	for name, check := range registry {
		go func(n string, c CheckFunc) {
			ch <- runWithTimeout(n, c)
		}(name, check)
	}

	var results []model.Result

	// collect results
	for range registry {
		results = append(results, <-ch)
	}

	return results
}

// RunOne executes a single check with timeout protection
func RunOne(name string) (model.Result, bool) {
	check, ok := registry[name]
	if !ok {
		return model.Result{}, false
	}

	return runWithTimeout(name, check), true
}

// --- core timeout wrapper ---

func runWithTimeout(name string, check CheckFunc) model.Result {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resultCh := make(chan model.Result, 1)

	go func() {
		resultCh <- check()
	}()

	select {
	case res := <-resultCh:
		return res

	case <-ctx.Done():
		return model.Result{
			Name:    name,
			Status:  model.Warning,
			Message: "timed out (2s)",
		}
	}
}
