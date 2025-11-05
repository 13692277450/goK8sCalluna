package utils

import (
	"regexp"
	"strings"
)

// KubeLogEntry represents a parsed Kubernetes log entry
type KubeLogEntry struct {
	Timestamp  string `json:"timestamp"`
	Node       string `json:"node"`
	Component  string `json:"component"`
	Level      string `json:"level"`
	DetailTime string `json:"detail_time"`
	PID        string `json:"pid"`
	FileLine   string `json:"file_line"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

// ParseKubeLog parses Kubernetes log format into structured data
func ParseKubeLog(logText string) ([]KubeLogEntry, error) {
	var entries []KubeLogEntry

	// Regex pattern for parsing Kubernetes logs
	// Example: Nov 02 09:20:22 k8s-master01 kubelet[2004]: E1102 09:20:22.223705    2004 cri_stats_provider.go:694] "Unable to fetch container log stats" err="failed to get fsstats for \"/var/log/pods/kube-system_calico-kube-controllers-558d465845-fglk8_4f55a4f1-035a-4a74-950d-4c51a56d1196/calico-kube-controllers/12.log\": no such file or directory" containerName="calico-kube-controllers"
	pattern := `(\w{3} \d{2} \d{2}:\d{2}:\d{2}) (\S+) (\S+)\[(\d+)\]: (\w)(\d{4} \d{2}:\d{2}:\d{2}\.\d{6}) +(\d+) (\S+)] "([^"]+)" err="([^"]+)"`
	re := regexp.MustCompile(pattern)

	lines := strings.SplitSeq(strings.TrimSpace(logText), "\n")

	for line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) >= 10 {
			entry := KubeLogEntry{
				Timestamp:  matches[1],
				Node:       matches[2],
				Component:  matches[3],
				Level:      matches[5],
				DetailTime: matches[6],
				PID:        matches[7],
				FileLine:   matches[8],
				Message:    matches[9],
				Error:      matches[10],
			}
			entries = append(entries, entry)
		}
	}

	return entries, nil
}

// LogToJSON converts log text to JSON string
func LogToJSON(logText string) (string, error) {
	entries, err := ParseKubeLog(logText)
	if err != nil {
		return "", err
	}

	// Convert to JSON - you would typically use json.Marshal here
	// For simplicity, returning as formatted string
	var result strings.Builder
	result.WriteString("[\n")
	for i, entry := range entries {
		if i > 0 {
			result.WriteString(",\n")
		}
		result.WriteString("  {\n")
		result.WriteString(`    "timestamp": "` + entry.Timestamp + `",\n`)
		result.WriteString(`    "node": "` + entry.Node + `",\n`)
		result.WriteString(`    "component": "` + entry.Component + `",\n`)
		result.WriteString(`    "level": "` + entry.Level + `",\n`)
		result.WriteString(`    "detail_time": "` + entry.DetailTime + `",\n`)
		result.WriteString(`    "pid": "` + entry.PID + `",\n`)
		result.WriteString(`    "file_line": "` + entry.FileLine + `",\n`)
		result.WriteString(`    "message": "` + entry.Message + `",\n`)
		result.WriteString(`    "error": "` + entry.Error + `"` + "\n")
		result.WriteString("  }")
	}
	result.WriteString("\n]")

	return result.String(), nil
}
