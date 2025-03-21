package log

import "time"

type LogEntry struct {
	TimeStamp       time.Time
	ProductId       string
	Destination     string
	CompleteLogLine string
}
