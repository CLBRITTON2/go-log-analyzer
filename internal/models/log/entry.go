package log

import "time"

type LogEntry struct {
	Timestamp   time.Time
	ProductId   string
	Destination string
}
