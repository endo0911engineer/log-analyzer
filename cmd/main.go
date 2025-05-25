package main

import (
	"flag"
	"fmt"
	"log-analyzer/internal/aggregator"
	"log-analyzer/internal/output"
	"log-analyzer/internal/parser"
	"os"
)

func main() {
	filepath := flag.String("file", "", "Path to the log file (JSON lines)")
	flag.Parse()

	if *filepath == "" {
		fmt.Println("Usage: log-analyzer -file <path>")
		os.Exit(1)
	}

	entries := parser.ParseLogFile(*filepath)
	stats := aggregator.Aggregate(entries)
	output.PrintJSON(stats)
}
