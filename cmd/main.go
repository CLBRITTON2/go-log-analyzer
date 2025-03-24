package main

import (
	"fmt"

	"github.com/CLBRITTON2/go-log-analyzer/internal/analyzer"
	"github.com/CLBRITTON2/go-log-analyzer/internal/parser"
)

func main() {
	// Just for testing at the moment
	coreLogFilePath := "../test/sample_logs/Core_20250323.log"
	coreLogEntries, coreParseErr := parser.ParseCoreLogFile(coreLogFilePath)
	if coreParseErr != nil {
		fmt.Printf("Error calling ParseCoreLogFile: %v\n", coreParseErr)
	}

	coreReportData, analyzerErr := analyzer.AnalyzeCoreLogEntries(coreLogEntries)
	if analyzerErr != nil {
		fmt.Printf("Error calling AnalyzeCoreLogEntries: %v\n", analyzerErr)
	}

	appLogFilePath := "../test/sample_logs/ClothingScraper_20250323.log"
	appLogEntries, appParseErr := parser.ParseAppLogFile(appLogFilePath)
	if appParseErr != nil {
		fmt.Printf("Error calling ParseAppLogFile: %v\n", appParseErr)
	}

	appReportData, analyzerErr := analyzer.AnalyzeAppLogEntries(appLogEntries)
	if analyzerErr != nil {
		fmt.Printf("Error calling AnalyzeAppLogEntries: %v\n", analyzerErr)
	}

	coreReportData.PrintCoreReportSummary()
	appReportData.PrintAppReportSummary()

}
