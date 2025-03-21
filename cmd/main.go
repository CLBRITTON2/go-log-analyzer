package main

import (
	"fmt"

	"github.com/CLBRITTON2/go-log-analyzer/internal/analyzer"
	"github.com/CLBRITTON2/go-log-analyzer/internal/parser"
)

func main() {
	// Just for testing at the moment
	logFilePath := "../test/sample_logs/Core_20250319.log"
	logEntries, parseErr := parser.ParseLogFile(logFilePath)
	if parseErr != nil {
		fmt.Printf("Error calling ParseLogFile: %v\n", parseErr)
	}

	logReportData, analyzerErr := analyzer.AnalyzeLogEntries(logEntries)
	if analyzerErr != nil {
		fmt.Printf("Error calling AnalyzeLogEntries: %v\n", analyzerErr)
	}

	fmt.Println("Log summary:")
	fmt.Printf("Total Products Processed: %v\n", logReportData.TotalProductsProcessed)
	fmt.Printf("Total Errors: %v\n", logReportData.ErrorCount)
	fmt.Printf("Total Blocked Products: %v\n", logReportData.BlockedProductCount)
	fmt.Printf("Total Submissions: %v\n", logReportData.SubmittedProductCount)
}
