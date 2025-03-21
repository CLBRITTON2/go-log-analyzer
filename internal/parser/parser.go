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
		fmt.Printf("ParseLogFile - Failed to open log file")
	}
	defer logFile.Close()

	var logEntries []log.LogEntry

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
				// If theres no timestamp skip the line - this is not an error some lines just dont get stamped
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
			logEntry := log.LogEntry{
				Timestamp:   timeStamp,
				ProductId:   productId,
				Destination: destination,
			}
			logEntries = append(logEntries, logEntry)
		}
	}

	// If there's an error scanning don't return log entries return nil
	if err := logScanner.Err(); err != nil {
		fmt.Printf("ParseLogFile - Error while scanning log file")
		return nil, err
	}

	return logEntries, nil
}
