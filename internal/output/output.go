package output

import (
	"encoding/json"
	"fmt"
	"log-analyzer/internal/aggregator"
)

func PrintJSON(stats aggregator.Stats) {
	out, err := json.MarshalIndent(stats, "", " ")
	if err != nil {
		fmt.Println("Failed to marshal output:", err)
		return
	}
	fmt.Println(string(out))
}
