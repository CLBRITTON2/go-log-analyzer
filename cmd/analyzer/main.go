package main

import (
	"fmt"

	"github.com/CLBRITTON2/go-log-analyzer/internal/parser"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Just for testing at the moment
	logFilePath := "../../test/sample_logs/Core_20250319.log"
	logEntries, err := parser.ParseLogFile(logFilePath)
	checkError(err)
	fmt.Printf("Found %d product entries\n", len(logEntries))
	// for i := 0; i < len(logEntries)/1000; i++ {
	// 	fmt.Println(logEntries[i].Timestamp.Format("2006-01-02 15:04:05"), logEntries[i].ProductId, logEntries[i].Destination)
	// }
}
