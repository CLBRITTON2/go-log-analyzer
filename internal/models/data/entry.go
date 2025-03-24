package data

import "time"

type CoreLogEntry struct {
	TimeStamp       time.Time
	ProductId       *string
	Destination     *string
	CompleteLogLine *string
}

type AppLogEntry struct {
	ApplicationName string
	CompletedCycles int
	CycleDuration   float64
}

type CoreLogEntries []CoreLogEntry

type AppLogEntries []AppLogEntry

func (coreLogEntries *CoreLogEntries) AddCoreLogEntry(timeStamp time.Time, productId *string, destination *string, completeLogLine *string) {
	CoreLogEntry := CoreLogEntry{
		TimeStamp:       timeStamp,
		ProductId:       productId,
		Destination:     destination,
		CompleteLogLine: completeLogLine,
	}
	*coreLogEntries = append(*coreLogEntries, CoreLogEntry)
}

func (appLogEntries *AppLogEntries) AddAppLogEntry(applicationName string, completedCycles int, cycleDuration float64) {
	appLogEntry := AppLogEntry{
		ApplicationName: applicationName,
		CompletedCycles: completedCycles,
		CycleDuration:   cycleDuration,
	}
	*appLogEntries = append(*appLogEntries, appLogEntry)
}
