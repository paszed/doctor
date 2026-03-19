package checks

import (
	"sort"
	"strings"
	"sync"

	"github.com/paszed/doctor/internal/model"
)

func RunAll() []model.Result {
	results := make(chan model.Result, len(registry))
	var wg sync.WaitGroup

	for _, check := range registry {
		wg.Add(1)

		go func(c CheckFunc) {
			defer wg.Done()
			results <- c()
		}(check)
	}

	wg.Wait()
	close(results)

	var all []model.Result

	for r := range results {
		all = append(all, r)
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i].Name < all[j].Name
	})

	return all
}

func RunOne(name string) (model.Result, bool) {
	fn, ok := registry[strings.ToLower(name)]
	if !ok {
		return model.Result{}, false
	}

	return fn(), true
}

func List() []string {
	var names []string

	for name := range registry {
		names = append(names, name)
	}

	sort.Strings(names)
	return names
}

func RunSelected(names []string) ([]model.Result, []string) {
	var results []model.Result
	var unknown []string

	for _, name := range names {
		fn, ok := registry[strings.ToLower(name)]
		if !ok {
			unknown = append(unknown, name)
			continue
		}

		results = append(results, fn())
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Name < results[j].Name
	})

	return results, unknown
}
