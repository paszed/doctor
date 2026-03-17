package state

import "github.com/paszed/doctor/internal/model"

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
