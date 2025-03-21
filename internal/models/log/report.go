package log

import (
	"os"
	"strconv"

	"github.com/aquasecurity/table"
)

type ReportData struct {
	TotalProductsProcessed int
	BlockedProductCount    int
	SubmittedProductCount  int
	ErrorCount             int
}

func (reportData ReportData) PrintReportSummary() {
	table := table.New(os.Stdout)
	table.AddRow(strconv.Itoa(reportData.TotalProductsProcessed), strconv.Itoa(reportData.BlockedProductCount), strconv.Itoa(reportData.SubmittedProductCount), strconv.Itoa(reportData.ErrorCount))
	table.SetHeaders("Total Products Processed", "Blocked Products", "Submitted Products", "Errors")
	table.Render()
}
