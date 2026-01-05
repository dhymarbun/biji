package main

import (
	"bufio"
	"os"
	"strings"
)

func AnalyzeLog(path string, baseline int) []TrafficAlert {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	counter := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) > 6 {
			counter[parts[6]]++
		}
	}

	var alerts []TrafficAlert
	for ep, count := range counter {
		if count > baseline {
			alerts = append(alerts, TrafficAlert{
				Endpoint: ep,
				Count:    count,
				Baseline: baseline,
				Risk:     CalculateTrafficRisk(count, baseline),
			})
		}
	}

	return alerts
}

func CalculateTrafficRisk(count, baseline int) string {
	ratio := count / baseline

	switch {
	case ratio >= 10:
		return "CRITICAL"
	case ratio >= 5:
		return "HIGH"
	case ratio >= 2:
		return "MEDIUM"
	default:
		return "LOW"
	}
}
