package ui

import (
	"encoding/json"
	"fmt"

	"github.com/paszed/doctor/internal/model"
)

func PrintJSON(results []model.Result) {

	out := make(map[string]string)

	for _, r := range results {
		out[r.Name] = r.Message
	}

	data, _ := json.MarshalIndent(out, "", "  ")
	fmt.Println(string(data))
}
