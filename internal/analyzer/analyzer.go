package analyzer

import (
	"strings"

	"github.com/CLBRITTON2/go-log-analyzer/internal/models/data"
)

func AnalyzeCoreLogEntries(coreLogEntries data.CoreLogEntries) (data.CoreReportData, error) {
	totalProductsProcessed, errorCount, blockedProductCount, submittedProductCount := 0, 0, 0, 0

	for i := range coreLogEntries {
		// No destination or product ID represents a generic log entry for errors/init/etc
		if coreLogEntries[i].ProductId == nil || coreLogEntries[i].Destination == nil {
			if strings.Contains(strings.ToLower(*coreLogEntries[i].CompleteLogLine), strings.ToLower("error")) || strings.Contains(strings.ToLower(*coreLogEntries[i].CompleteLogLine), strings.ToLower("failed")) {
				errorCount++
			}
		} else {
			totalProductsProcessed++
			if strings.Contains(strings.ToLower(*coreLogEntries[i].Destination), strings.ToLower("blocklist")) {
				blockedProductCount++
			}
			if strings.Contains(strings.ToLower(*coreLogEntries[i].Destination), strings.ToLower("discord")) {
				submittedProductCount++
			}
		}
	}

	logReportData := data.CoreReportData{
		TotalProductsProcessed: totalProductsProcessed,
		BlockedProductCount:    blockedProductCount,
		SubmittedProductCount:  submittedProductCount,
		ErrorCount:             errorCount,
	}

	return logReportData, nil
}

func AnalyzeAppLogEntries(appLogEntries data.AppLogEntries) (data.AppReportData, error) {
	totalCyclesCompleted := 0
	averageCycleDuration := 0.0
	// Every entry has the app name included so index doesn't matter here
	applicationName := appLogEntries[0].ApplicationName
	cycleDurations := make([]float64, len(appLogEntries))

	for i := range appLogEntries {
		if appLogEntries[i].CycleDuration != 0.0 {
			totalCyclesCompleted++
			cycleDurations[i] = appLogEntries[i].CycleDuration
		}
	}

	averageCycleDuration = CalculateAverageCycleDuration(cycleDurations)

	logReportData := data.AppReportData{
		ApplicationName:      applicationName,
		TotalCyclesCompleted: totalCyclesCompleted,
		AverageCycleDuration: averageCycleDuration,
	}

	return logReportData, nil
}

func CalculateAverageCycleDuration(cycleDurations []float64) float64 {
	totalDuration := 0.0
	for i := 0; i < len(cycleDurations); i++ {
		totalDuration += cycleDurations[i]
	}
	return totalDuration / float64(len(cycleDurations))
}
