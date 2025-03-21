package log

import "time"

type LogEntry struct {
	TimeStamp       time.Time
	ProductId       *string
	Destination     *string
	CompleteLogLine *string
}

type LogEntries []LogEntry

func (logEntries *LogEntries) AddLogEntry(timeStamp time.Time, productId *string, destination *string, completeLogLine *string) {
	logEntry := LogEntry{
		TimeStamp:       timeStamp,
		ProductId:       productId,
		Destination:     destination,
		CompleteLogLine: completeLogLine,
	}
	*logEntries = append(*logEntries, logEntry)
}
