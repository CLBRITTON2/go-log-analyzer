package main

import (
	"flag"
)

type CLIFlags struct {
	ShowCoreLogSummary    bool
	ShowAllAppLogSummary  bool
	TotalProductsOnly     bool
	BlockedProductsOnly   bool
	SubmittedProductsOnly bool
	ErrorsOnly            bool
}

func CreateCLIFlags() *CLIFlags {
	flags := &CLIFlags{}

	flag.BoolVar(&flags.ShowCoreLogSummary, "core", true, "Show summary of the core log file only")
	flag.BoolVar(&flags.ShowAllAppLogSummary, "apps", false, "Show summary of all app log files including core")
	flag.BoolVar(&flags.TotalProductsOnly, "products", false, "Show only total products processed from core")
	flag.BoolVar(&flags.BlockedProductsOnly, "blocked", false, "Show only blocked products from core")
	flag.BoolVar(&flags.SubmittedProductsOnly, "submitted", false, "Show only submitted products from core")
	flag.BoolVar(&flags.ErrorsOnly, "errors", false, "Show only errors from core")

	return flags
}
