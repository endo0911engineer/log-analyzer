package parser

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type LogEntry struct {
	Timestamp        string `json:"timestamp"`
	UserID           string `json:"user_id"`
	Model            string `json:"model"`
	PromptTokens     int    `json:"prompt_tokens"`
	CompletionTokens int    `json:"completion_tokens"`
	StatusCode       int    `json:"status_code"`
	LatencyMs        int    `json:"latency_ms"`
}

func ParseLogFile(path string) ([]LogEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	var entries []LogEntry
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var entry LogEntry
		if err := json.Unmarshal(scanner.Bytes(), &entry); err == nil {
			entries = append(entries, entry)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}
