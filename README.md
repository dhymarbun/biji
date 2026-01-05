<img width="841" height="747" alt="Screenshot 2026-01-03 123230" src="https://github.com/user-attachments/assets/26ca57b5-59b4-4e38-8559-a6abcf9ab010" />

# Biji

**Biji** is a defensive cybersecurity command-line tool designed for **passive web risk analysis** and **traffic anomaly detection**.

This project focuses on identifying potential security misconfigurations and abnormal traffic patterns **without performing any intrusive or exploitative actions**.  
Biji was built as a first cybersecurity portfolio project, emphasizing **ethical security practices**, **clear risk communication**, and a **blue-team mindset**.

---

## Features

### Passive Web Risk Analysis
- Detects missing HTTP security headers based on OWASP Secure Headers guidance
- Calculates a **risk score** and **risk severity level**
- Provides **recommended defensive actions** for each missing header
- Performs **read-only HTTP requests** (no payloads, no exploitation)

### Traffic Anomaly Detection (Log-Based)
- Analyzes web server access logs
- Detects abnormal request patterns using baseline thresholds
- Classifies traffic severity (LOW, MEDIUM, HIGH, CRITICAL)
- Does **not** generate or simulate traffic

### Security Report Mode
- Export scan results to a readable report file
- Suitable for documentation and learning purposes

### CLI Experience
- Clean, colorized terminal output
- Beginner-friendly `--help` command
- Tool metadata via `--version`

---

## What Biji Does NOT Do

- No exploitation or vulnerability confirmation
- No brute-force or fuzzing
- No traffic generation or stress testing
- No bypass or evasion techniques

All findings represent **potential security risks**, not confirmed vulnerabilities.

---

## Installation

### Prerequisites
- Go 1.22 or newer

### Build
```bash
go build -o biji
```

## Usage
Show Help
```bash
./biji --help
```

Show Version
```bash
./biji --version
```

Passive Web Risk Scan
```bash
./biji --scan https://example.com
```

Web Scan + Report
```bash
./biji --scan https://example.com --report report.txt
```
Traffic Anomaly Detection
```bash
./biji --monitor access.log --threshold 100
```
