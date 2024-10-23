package main

import (
	"fmt"
	"practice_1/alerts"
	"practice_1/metrics"
	"practice_1/polling"
)

const serverURL = "http://srv.msk01.gigacorp.local/_stats"
const maxRetryCount = 3

func main() {
	pollServer := polling.InitiatePolling(serverURL, maxRetryCount)

	for response := range pollServer() {
		serverMetrics, err := metrics.ParseMetrics(response)
		if err != nil {
			fmt.Printf("Error parsing metrics: %v\n", err)
			continue
		}

		alerts.CheckMetrics(serverMetrics)
	}
}
