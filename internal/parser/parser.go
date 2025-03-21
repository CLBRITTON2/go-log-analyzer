package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/CLBRITTON2/go-log-analyzer/internal/models/log"
)

var timestampPattern = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`)
var productIdPattern = regexp.MustCompile(`ASIN\s([0-9A-Z]{10})`)
var destinationPattern = regexp.MustCompile(`to\s([\w/]+\.txt)`)

func ParseLogFile(filePath string) ([]log.LogEntry, error) {
	logFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	logEntries := log.LogEntries{}

	logScanner := bufio.NewScanner(logFile)
	for logScanner.Scan() {
		line := logScanner.Text()

		// Element 0 is the full matched string (line from the log file) - element 1 is the captured regex match for all regex patterns
		timeStampMatches := timestampPattern.FindStringSubmatch(line)
		var timeStamp time.Time
		var timeStampParsingError error
		if len(timeStampMatches) > 1 {
			timeStamp, timeStampParsingError = time.Parse("2006-01-02 15:04:05", timeStampMatches[1])
			if timeStampParsingError != nil {
				// If theres no timestamp skip the line
				continue
			}
		}

		productIdMatch := productIdPattern.FindStringSubmatch(line)
		var productId string
		if len(productIdMatch) > 1 {
			productId = productIdMatch[1]
		}

		destinationMatch := destinationPattern.FindStringSubmatch(line)
		var destination string
		if len(destinationMatch) > 1 {
			destination = destinationMatch[1]
		}

		if !timeStamp.IsZero() && productId != "" && destination != "" {
			logEntries.AddLogEntry(timeStamp, &productId, &destination, nil)
		} else {
			// Storing the full line for entries that aren't product related - these could be errors
			// and need to be captured for reporting
			logEntries.AddLogEntry(timeStamp, nil, nil, &line)
		}
	}

	if err := logScanner.Err(); err != nil {
		return nil, fmt.Errorf("Error while scanning log file: %v", err)
	}

	return logEntries, nil
}
