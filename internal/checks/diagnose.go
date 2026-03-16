package checks

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/paszed/doctor/internal/model"
	"github.com/paszed/doctor/internal/ui"
)

func RunDiagnose() {

	start := time.Now()

	ui.Section("TOOLS")

	results := make(chan model.Result, len(Registry))
	var wg sync.WaitGroup

	for _, check := range Registry {
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

	for _, r := range all {
		ui.PrintResult(r)
	}

	ok := 0
	missing := 0
	warn := 0

	for _, r := range all {
		switch r.Status {
		case model.OK:
			ok++
		case model.Missing:
			missing++
		case model.Warning:
			warn++
		}
	}

	ui.Section("SUMMARY")

	fmt.Printf("✓ %d/%d checks passed\n", ok, len(all))

	if warn > 0 {
		fmt.Printf("! %d warnings\n", warn)
	}

	if missing > 0 {
		fmt.Printf("✗ %d missing\n", missing)
	}

	duration := time.Since(start)
	fmt.Printf("\ncompleted in %s\n", duration.Round(time.Millisecond))
}
