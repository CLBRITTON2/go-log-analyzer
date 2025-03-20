package parser

import (
	"bufio"
	"os"
	"regexp"
	"time"

	"github.com/CLBRITTON2/go-log-analyzer/internal/models/log"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseLogFile(filePath string) ([]log.LogEntry, error) {
	logFile, err := os.Open(filePath)
	checkError(err)
	defer logFile.Close()

	timestampPattern := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`)
	var logEntries []log.LogEntry

	logScanner := bufio.NewScanner(logFile)
	for logScanner.Scan() {
		line := logScanner.Text()
		timestampMatches := timestampPattern.FindStringSubmatch(line)
		// Skip lines without a timestamp match - this isnt an error some logs could be corrupt
		if len(timestampMatches) < 2 {
			continue
		}

		// Element 0 is the full matched string - element 1 is the captured timestamp
		timeStamp, err := time.Parse("2006-01-02 15:04:05", timestampMatches[1])
		checkError(err)

		logEntry := log.LogEntry{
			Timestamp: timeStamp,
		}
		logEntries = append(logEntries, logEntry)
	}
	return logEntries, nil
}
