package main

import (
	"net/http"
	"time"
)

func ScanWeb(url string) WebRiskResult {
	client := http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return WebRiskResult{
			URL:       url,
			RiskLevel: "UNKNOWN",
		}
	}
	defer resp.Body.Close()

	required := []string{
		"Content-Security-Policy",
		"Strict-Transport-Security",
		"X-Frame-Options",
		"X-Content-Type-Options",
	}

	var missing []string
	for _, h := range required {
		if resp.Header.Get(h) == "" {
			missing = append(missing, h)
		}
	}

	score, level := CalculateWebRisk(missing, resp.Header.Get("Server"))

	return WebRiskResult{
		URL:            url,
		MissingHeaders: missing,
		ServerHeader:   resp.Header.Get("Server"),
		RiskScore:      score,
		RiskLevel:      level,
	}
}

func CalculateWebRisk(missing []string, server string) (int, string) {
	score := 0

	for _, h := range missing {
		switch h {
		case "Content-Security-Policy":
			score += 30
		case "Strict-Transport-Security":
			score += 25
		case "X-Frame-Options":
			score += 15
		case "X-Content-Type-Options":
			score += 10
		}
	}

	if server != "" {
		score += 10
	}

	level := "LOW"
	switch {
	case score >= 76:
		level = "CRITICAL"
	case score >= 51:
		level = "HIGH"
	case score >= 26:
		level = "MEDIUM"
	}

	return score, level
}
