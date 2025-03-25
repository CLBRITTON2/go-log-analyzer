package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/CLBRITTON2/go-log-analyzer/internal/analyzer"
	"github.com/CLBRITTON2/go-log-analyzer/internal/parser"
)

func main() {
	flags := CreateCLIFlags()
	flag.Parse()
	// For testing
	// logDirectory := "../test/sample_logs"
	executingPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	executingDirectory := filepath.Dir(executingPath)

	if flags.ShowCoreLogSummary {
		coreLogFile, findCoreLogErr := findCoreLogFile(executingDirectory)
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
		appLogFiles, findAppLogErr := findAppLogFiles(executingDirectory)
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

func findCoreLogFile(rootDirectory string) (string, error) {
	var coreLogFilePath string

	err := filepath.WalkDir(rootDirectory, func(path string, directory fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(strings.ToLower(directory.Name()), "core") && strings.HasSuffix(directory.Name(), ".log") {
			coreLogFilePath = path
			return fs.SkipDir
		}
		return nil
	})

	if err != nil {
		return "", fmt.Errorf("error walking directory: %w", err)
	}

	if coreLogFilePath == "" {
		return "", fmt.Errorf("no core log file found")
	}

	return coreLogFilePath, nil
}

func findAppLogFiles(rootDirectory string) ([]string, error) {
	var appLogFilePaths []string
	err := filepath.WalkDir(rootDirectory, func(path string, directory fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(strings.ToLower(directory.Name()), "scraper") && strings.HasSuffix(directory.Name(), ".log") {
			appLogFilePaths = append(appLogFilePaths, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error walking through directory: %v\n", err)
	}
	if len(appLogFilePaths) == 0 {
		return nil, fmt.Errorf("No app log files found")
	}

	return appLogFilePaths, nil
}
