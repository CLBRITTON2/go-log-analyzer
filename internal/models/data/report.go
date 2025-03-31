package data

import (
	"fmt"
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
	table.SetHeaders("Total Products", "Blocked", "Submitted", "Errors")
	table.Render()
}

func (reportData CoreReportData) PrintTotalProductsProcessed() {
	fmt.Printf("Core Total Products Processed: %d\n", reportData.TotalProductsProcessed)
}

func (reportData CoreReportData) PrintBlockedProductCount() {
	fmt.Printf("Core Blocked Product Count: %d\n", reportData.BlockedProductCount)
}

func (reportData CoreReportData) PrintSubmittedProductCount() {
	fmt.Printf("Core Submitted Product Count: %d\n", reportData.SubmittedProductCount)
}

func (reportData CoreReportData) PrintErrorCount() {
	fmt.Printf("Core Error Count: %d\n", reportData.ErrorCount)
}

func (appReportData AppReportData) PrintAppReportSummary() {
	table := table.New(os.Stdout)
	table.AddRow(appReportData.ApplicationName, strconv.Itoa(appReportData.TotalCyclesCompleted), strconv.FormatFloat(appReportData.AverageCycleDuration, 'f', 2, 64))
	table.SetHeaders("Application", "Total Cycles", "Average Cycle (sec)")
	table.Render()
}
