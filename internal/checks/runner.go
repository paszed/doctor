package checks

import (
	"fmt"
	"sync"
	"time"

	"github.com/paszed/doctor/internal/config"
	"github.com/paszed/doctor/internal/detect"
	"github.com/paszed/doctor/internal/model"
)

type resultWrapper struct {
	name   string
	result model.Result
}

func RunAll() []model.Result {
	var wg sync.WaitGroup
	resultsChan := make(chan resultWrapper, len(registry))

	// 🔥 project detection
	project := detect.Detect()

	for name, check := range registry {

		// 🔥 skip ignored checks
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

	// 🔥 DEDUPLICATE + ADD PORTS
	seenPorts := make(map[string]bool)

	for _, r := range results {
		if len(r.Name) > 5 && r.Name[:5] == "port:" {
			seenPorts[r.Name] = true
		}
	}

	// config ports
	for _, port := range config.Current.Ports {
		name := fmt.Sprintf("port:%d", port)

		if seenPorts[name] {
			continue
		}

		results = append(results, CheckPort(fmt.Sprintf("%d", port)))
		seenPorts[name] = true
	}

	// 🔥 auto-detect ports (Node project)
	if project.HasNode {
		name := "port:3000"
		if !seenPorts[name] {
			results = append(results, CheckPort("3000"))
		}
	}

	return results
}

// helper
func contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}
