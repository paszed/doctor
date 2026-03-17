package checks

import (
	"sort"
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
