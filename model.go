package main

type WebRiskResult struct {
	URL            string
	MissingHeaders []string
	ServerHeader   string
	RiskScore      int
	RiskLevel      string
}

type TrafficAlert struct {
	Endpoint string
	Count    int
	Baseline int
	Risk     string
}
