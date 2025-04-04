package parser

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/CLBRITTON2/go-log-analyzer/internal/models/data"
)

var (
	timestampPattern   = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`)
	asinPattern        = regexp.MustCompile(`ASIN\s([0-9A-Z]{10})`)
	upcPattern         = regexp.MustCompile(`UPC\s(\d+)`)
	destinationPattern = regexp.MustCompile(`to\s([\w/]+\.txt|discord)`)
	durationPattern    = regexp.MustCompile(`completed in (\d+\.\d+) seconds`)
)

func ParseCoreLogFile(filePath string) (data.CoreLogEntries, error) {
	coreLogFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer coreLogFile.Close()

	coreLogEntries := data.CoreLogEntries{}

	coreLogScanner := bufio.NewScanner(coreLogFile)
	for coreLogScanner.Scan() {
		line := coreLogScanner.Text()

		// Element 0 is the full matched string (line from the log file) - element 1 is the captured regex match for all regex patterns
		timeStampMatches := timestampPattern.FindStringSubmatch(line)
		var timeStamp time.Time
		if len(timeStampMatches) > 1 {
			timeStamp, err = time.Parse("2006-01-02 15:04:05", timeStampMatches[1])
			if err != nil {
				continue
			}
		}

		// Check for Amazon products first
		asinMatch := asinPattern.FindStringSubmatch(line)
		var productId string
		if len(asinMatch) > 1 {
			productId = asinMatch[1]
		}

		// Check for products that have a UPC - a product wont log a upc and ASIN so this is safe
		upcMatch := upcPattern.FindStringSubmatch(line)
		if len(upcMatch) > 1 {
			productId = upcMatch[1]
		}

		destinationMatch := destinationPattern.FindStringSubmatch(line)
		var destination string
		if len(destinationMatch) > 1 {
			destination = destinationMatch[1]
		}

		if productId != "" && destination != "" {
			coreLogEntries.AddCoreLogEntry(timeStamp, &productId, &destination, nil)
		} else {
			// Storing the full line for entries that aren't product related - these could be errors
			// and need to be captured for reporting
			coreLogEntries.AddCoreLogEntry(timeStamp, nil, nil, &line)
		}
	}

	if err := coreLogScanner.Err(); err != nil {
		return nil, fmt.Errorf("error while scanning log file: %v", err)
	}

	return coreLogEntries, nil
}

func ParseAppLogFile(filePath string) (data.AppLogEntries, error) {
	appLogFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer appLogFile.Close()

	// Use the file name to pull the app name removing app log suffix without using REGEX
	// Removing the suffix isn't necessary but it makes the report look cleaner

	filename := filepath.Base(filePath)
	appSuffix := "Scraper_"
	suffixIndex := strings.Index(filename, appSuffix)
	var applicationName string
	if suffixIndex > 0 {
		applicationName = filename[:suffixIndex]
	} else {
		applicationName = "Unknown"
	}

	appLogEntries := data.AppLogEntries{}
	cycleDuration := 0.0

	appLogScanner := bufio.NewScanner(appLogFile)
	for appLogScanner.Scan() {
		line := appLogScanner.Text()

		completedCycles := 0
		if strings.Contains(strings.ToLower(line), strings.ToLower("run completed")) {
			completedCycles = 1
		}

		cycleDurationMatches := durationPattern.FindStringSubmatch(line)
		if len(cycleDurationMatches) > 1 {
			cycleDuration, err = strconv.ParseFloat(cycleDurationMatches[1], 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse cycle duration: %v", err)
			}
		}

		if cycleDuration > 0.0 && completedCycles == 1 {
			appLogEntries.AddAppLogEntry(applicationName, completedCycles, cycleDuration)
		}
	}

	if err := appLogScanner.Err(); err != nil {
		return nil, fmt.Errorf("error while scanning log file: %v", err)
	}

	return appLogEntries, nil
}

func FindCoreLogFile(rootDirectory string) (string, error) {
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

func FindAppLogFiles(rootDirectory string) ([]string, error) {
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
		return nil, fmt.Errorf("error walking through directory: %v", err)
	}
	if len(appLogFilePaths) == 0 {
		return nil, fmt.Errorf("no app log files found")
	}

	return appLogFilePaths, nil
}
