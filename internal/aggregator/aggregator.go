package aggregator

import "log-analyzer/internal/parser"

type Stats struct {
	TotalRequests    int            `json:"total_requests"`
	StatusCodeCounts map[int]int    `json:"status_code_counts"`
	ModelCounts      map[string]int `json:"model_counts"`
	TotalTokens      int            `json:"total_tokens"`
	AverageLatencyMs float64        `json:"average_latency_ms"`
}

func Aggregate(entries []parser.LogEntry) Stats {
	var stats Stats
	stats.StatusCodeCounts = make(map[int]int)
	stats.ModelCounts = make(map[string]int)

	var totalLatency int

	for _, entry := range entries {
		stats.TotalRequests++
		stats.StatusCodeCounts[entry.StatusCode]++
		stats.ModelCounts[entry.Model]++
		stats.TotalTokens += entry.PromptTokens + entry.CompletionTokens
		totalLatency += entry.LatencyMs
	}

	if stats.TotalRequests > 0 {
		stats.AverageLatencyMs = float64(totalLatency) / float64(stats.TotalRequests)
	}

	return stats
}
