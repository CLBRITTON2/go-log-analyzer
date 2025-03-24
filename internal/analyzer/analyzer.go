package analyzer

import (
	"strings"

	"github.com/CLBRITTON2/go-log-analyzer/internal/models/data"
)

func AnalyzeLogEntries(logEntries data.LogEntries) (data.ReportData, error) {
	totalProductsProcessed, errorCount, blockedProductCount, submittedProductCount := 0, 0, 0, 0

	for i := 0; i < len(logEntries); i++ {
		// No destination or product ID represents a generic log entry for errors/init/etc
		if logEntries[i].ProductId == nil || logEntries[i].Destination == nil {
			if strings.Contains(strings.ToLower(*logEntries[i].CompleteLogLine), strings.ToLower("error")) || strings.Contains(strings.ToLower(*logEntries[i].CompleteLogLine), strings.ToLower("failed")) {
				errorCount++
			}
		} else {
			totalProductsProcessed++
			if strings.Contains(strings.ToLower(*logEntries[i].Destination), strings.ToLower("blocklist")) {
				blockedProductCount++
			}
			if strings.Contains(strings.ToLower(*logEntries[i].Destination), strings.ToLower("discord")) {
				submittedProductCount++
			}
		}
	}

	logReportData := data.ReportData{
		TotalProductsProcessed: totalProductsProcessed,
		BlockedProductCount:    blockedProductCount,
		SubmittedProductCount:  submittedProductCount,
		ErrorCount:             errorCount,
	}

	return logReportData, nil
}
