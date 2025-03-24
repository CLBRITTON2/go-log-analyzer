package data

import (
	"os"
	"strconv"

	"github.com/aquasecurity/table"
)

type CoreReportData struct {
	TotalProductsProcessed int
	BlockedProductCount    int
	SubmittedProductCount  int
	ErrorCount             int
}

type AppReportData struct {
	ApplicationName      string
	TotalCyclesCompleted int
	AverageCycleDuration float64
}

func (reportData CoreReportData) PrintCoreReportSummary() {
	table := table.New(os.Stdout)
	table.AddRow(strconv.Itoa(reportData.TotalProductsProcessed), strconv.Itoa(reportData.BlockedProductCount), strconv.Itoa(reportData.SubmittedProductCount), strconv.Itoa(reportData.ErrorCount))
	table.SetHeaders("Total Products", "Blocked Products", "Submitted Products", "Errors")
	table.Render()
}

func (appReportData AppReportData) PrintAppReportSummary() {
	table := table.New(os.Stdout)
	table.AddRow(appReportData.ApplicationName, strconv.Itoa(appReportData.TotalCyclesCompleted), strconv.FormatFloat(appReportData.AverageCycleDuration, 'f', 2, 64))
	table.SetHeaders("Application", "Total Cycles", "Average Cycle Duration (sec)")
	table.Render()
}
