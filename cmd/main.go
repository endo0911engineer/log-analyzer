package main

import (
	"flag"
	"fmt"
	"log"
	"log-analyzer/internal/aggregator"
	"log-analyzer/internal/output"
	"log-analyzer/internal/parser"
	"strings"
	"sync"
)

func main() {
	filesFlag := flag.String("file", "", "Comma separated list of log files")
	flag.Parse()

	if *filesFlag == "" {
		fmt.Println("Usage: log-analyzer -file1.jsonl, file2.jsonl")
		return
	}

	files := strings.Split(*filesFlag, ",")

	var wg sync.WaitGroup
	entriesCh := make(chan []parser.LogEntry, len(files))
	errCh := make(chan error, len(files))

	for _, file := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			entries, err := parser.ParseLogFile(f)
			if err != nil {
				errCh <- fmt.Errorf("failed to parse %s: %w", f, err)
				return
			}
			entriesCh <- entries
		}(file)
	}

	wg.Wait()
	close(entriesCh)
	close(errCh)

	for err := range errCh {
		log.Fatal(err)
	}

	// すべてのエントリーをまとめる
	var allEntries []parser.LogEntry
	for entries := range entriesCh {
		allEntries = append(allEntries, entries...)
	}

	stats := aggregator.Aggregate(allEntries)
	output.PrintJSON(stats)
}
