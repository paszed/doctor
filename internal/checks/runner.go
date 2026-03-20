package checks

import (
	"fmt"
	"sync"
	"time"

	"github.com/paszed/doctor/internal/config"
	"github.com/paszed/doctor/internal/model"
)

type resultWrapper struct {
	name   string
	result model.Result
}

func RunAll() []model.Result {
	var wg sync.WaitGroup
	resultsChan := make(chan resultWrapper, len(registry))

	for name, check := range registry {

		// 🔥 Skip ignored checks
		if contains(config.Current.Ignore, name) {
			continue
		}

		wg.Add(1)

		go func(name string, check CheckFunc) {
			defer wg.Done()

			done := make(chan model.Result, 1)

			go func() {
				done <- check()
			}()

			select {
			case res := <-done:
				resultsChan <- resultWrapper{name: name, result: res}

			case <-time.After(2 * time.Second):
				resultsChan <- resultWrapper{
					name: name,
					result: model.Result{
						Name:    name,
						Status:  model.Warning,
						Message: "timed out (2s)",
					},
				}
			}
		}(name, check)
	}

	wg.Wait()
	close(resultsChan)

	var results []model.Result

	for r := range resultsChan {
		results = append(results, r.result)
	}

	// 🔥 Add dynamic port checks from config
	for _, port := range config.Current.Ports {
		results = append(results, CheckPort(fmt.Sprintf("%d", port)))
	}

	return results
}

func contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}
