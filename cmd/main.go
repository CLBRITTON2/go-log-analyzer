package main

import (
	"flag"
	"fmt"

	"github.com/CLBRITTON2/go-log-analyzer/internal/analyzer"
	"github.com/CLBRITTON2/go-log-analyzer/internal/parser"
)

func main() {
	flags := CreateCLIFlags()
	flag.Parse()
	// For testing
	logDirectory := "../test/sample_logs"
	// executingPath, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }
	// executingDirectory := filepath.Dir(executingPath)

	if flags.ShowCoreLogSummary {
		coreLogFile, findCoreLogErr := parser.FindCoreLogFile(logDirectory)
		if findCoreLogErr != nil {
			fmt.Printf("Error calling findCoreLogFile: %v\n", findCoreLogErr)
		}

		coreLogEntries, coreParseErr := parser.ParseCoreLogFile(coreLogFile)
		if coreParseErr != nil {
			fmt.Printf("Error calling ParseCoreLogFile: %v\n", coreParseErr)
		}

		coreReportData, analyzerErr := analyzer.AnalyzeCoreLogEntries(coreLogEntries)
		if analyzerErr != nil {
			fmt.Printf("Error calling AnalyzeCoreLogEntries: %v\n", analyzerErr)
		}

		switch {
		case flags.TotalProductsOnly:
			coreReportData.PrintTotalProductsProcessed()
		case flags.BlockedProductsOnly:
			coreReportData.PrintBlockedProductCount()
		case flags.SubmittedProductsOnly:
			coreReportData.PrintSubmittedProductCount()
		case flags.ErrorsOnly:
			coreReportData.PrintErrorCount()
		default:
			coreReportData.PrintCoreReportSummary()
		}
	}

	if flags.ShowAllAppLogSummary {
		appLogFiles, findAppLogErr := parser.FindAppLogFiles(logDirectory)
		if findAppLogErr != nil {
			fmt.Printf("Error calling findAppLogFiles: %v\n", findAppLogErr)
		}

		for i := range appLogFiles {
			appLogFilePath := appLogFiles[i]
			// fmt.Printf("Found app log file: %s\n", appLogFilePath)
			appLogEntries, appParseErr := parser.ParseAppLogFile(appLogFilePath)
			if appParseErr != nil {
				fmt.Printf("Error calling ParseAppLogFile: %v\n", appParseErr)
			}
			appReportData, analyzerErr := analyzer.AnalyzeAppLogEntries(appLogEntries)
			if analyzerErr != nil {
				fmt.Printf("Error calling AnalyzeAppLogEntries: %v\n", analyzerErr)
			}
			appReportData.PrintAppReportSummary()
		}
	}
}
