package state

import "github.com/paszed/doctor/internal/model"

// Summary aggregates results into counts
func Summary(results []model.Result) (ok, warn, missing int) {
	for _, r := range results {
		switch r.Status {
		case model.OK:
			ok++
		case model.Warning:
			warn++
		case model.Missing:
			missing++
		}
	}
	return
}

// Status returns overall status priority
func Status(results []model.Result) model.Status {
	ok, warn, missing := Summary(results)

	if missing > 0 {
		return model.Missing
	}
	if warn > 0 {
		return model.Warning
	}
	if ok > 0 {
		return model.OK
	}

	return model.OK
}
