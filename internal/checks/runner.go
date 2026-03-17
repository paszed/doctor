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

	name = strings.ToLower(name)

	for _, check := range registry {

		result := check()

		if strings.ToLower(result.Name) == name {
			return result, true
		}
	}

	return model.Result{}, false
}
