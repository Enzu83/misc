package main

import (
	"fmt"
	"strings"
)

// ExtractCheckInfo extracts the configuration information for a specific check
// from a structured string containing multiple check configurations.
//
// Parameters:
//   - configOutput: The full configuration output string
//   - checkName: The name of the check to extract (e.g., "process_agent", "telemetry")
//
// Returns:
//   - The complete section for the specified check, or empty string if not found
func ExtractCheckInfo(configOutput, checkName string) string {
	// Create the start and end patterns
	startPattern := fmt.Sprintf("=== %s check ===", checkName)
	
	// Find the start of the check section
	startIndex := strings.Index(configOutput, startPattern)
	if startIndex == -1 {
		return "" // Check not found
	}
	
	// Find the start of the next check section or end of string
	remainingText := configOutput[startIndex:]
	lines := strings.Split(remainingText, "\n")
	
	var result []string
	inSection := false
	
	for _, line := range lines {
		if strings.HasPrefix(line, "=== ") && strings.HasSuffix(line, " check ===") {
			if inSection {
				// We've reached the next check section, stop here
				break
			}
			if line == startPattern {
				// This is our target check, start collecting
				inSection = true
				result = append(result, line)
			}
		} else if inSection {
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}

// Alternative more concise implementation using regex-like approach
func ExtractCheckInfoSimple(configOutput, checkName string) string {
	startPattern := fmt.Sprintf("=== %s check ===", checkName)
	nextCheckPattern := "\n=== "
	
	// Find start position
	startPos := strings.Index(configOutput, startPattern)
	if startPos == -1 {
		return ""
	}
	
	// Find the position after our check section
	searchFrom := startPos + len(startPattern)
	nextCheckPos := strings.Index(configOutput[searchFrom:], nextCheckPattern)
	
	var endPos int
	if nextCheckPos == -1 {
		// This is the last check, take everything to the end
		endPos = len(configOutput)
	} else {
		// Take everything up to the next check
		endPos = searchFrom + nextCheckPos
	}
	
	result := configOutput[startPos:endPos]
	
	// Clean up trailing whitespace and separators
	result = strings.TrimRight(result, "\n ")
	
	return result
}

// Example usage
func main() {
	sampleInput := `=== io check ===
Configuration provider: file
Configuration source: file:/opt/datadog-agent/etc/conf.d/io.d/conf.yaml.default
Config for instance ID: io:541b60d158de04a7
{}
~
===

=== jamf check ===
Configuration provider: file
Configuration source: file:/opt/datadog-agent/etc/conf.d/jamf.d/conf.yaml
Log Config:
logs:
  - path: /private/var/log/jamf.log
    service: jamf_logs
    source: jamf
    type: file
===

=== load check ===
Configuration provider: file
Configuration source: file:/opt/datadog-agent/etc/conf.d/load.d/conf.yaml.default
Config for instance ID: load:bf7cea93fb3aa780
{}
~
===

=== memory check ===
Configuration provider: file
Configuration source: file:/opt/datadog-agent/etc/conf.d/memory.d/conf.yaml.default
Config for instance ID: memory:3f1f6288b95b9979
{}
~
===

=== munki check ===
Configuration provider: file
Configuration source: file:/opt/datadog-agent/etc/conf.d/munki.d/conf.yaml
Log Config:
logs:
  - log_processing_rules:
      - name: new_log_start_with_date
        pattern: ^### Starting managedsoftwareupdate$
        type: multi_line
    path: /Library/Managed Installs/Logs/ManagedSoftwareUpdate.log
    service: munki_logs
    source: munki_agent
    type: file
===

=== network check ===
Configuration provider: file
Configuration source: file:/opt/datadog-agent/etc/conf.d/network.d/conf.yaml.default
Config for instance ID: network:4b0649b7e11f0772
{}
~
===

=== ntp check ===
Configuration provider: file
Configuration source: file:/opt/datadog-agent/etc/conf.d/ntp.d/conf.yaml.default
Config for instance ID: ntp:3c427a42a70bbf8
{}
~
===

=== process_agent check ===
Configuration provider: file
Configuration source: file:/opt/datadog-agent/etc/conf.d/process_agent.yaml.default
Config for instance ID: process_agent:ec116abde0f285c4
{}
~
===

=== telemetry check ===
Configuration provider: file
Configuration source: file:/opt/datadog-agent/etc/conf.d/telemetry.d/conf.yaml.default
Config for instance ID: telemetry:4d459fc318a47aa4
{}
~
===

=== uptime check ===
Configuration provider: file
Configuration source: file:/opt/datadog-agent/etc/conf.d/uptime.d/conf.yaml.default
Config for instance ID: uptime:c72f390abdefdf1a
{}
~
===

=== vscode check ===
Configuration provider: file
Configuration source: file:/opt/datadog-agent/etc/conf.d/vscode.d/conf.yaml
Log Config:
logs:
  - path: /Users/maxime.chambre/Library/Application Support/Code/logs/*/*/exthost/vscode.typescript-language-features/TypeScript.log
    service: vscode_ts_logs
    source: vscode_client
    type: file
===`

	// Test the function
	fmt.Println("=== Testing ExtractCheckInfo ===")
	
	// Test existing check
	result := ExtractCheckInfo(sampleInput, "telemetry")
	fmt.Printf("Result for 'telemetry':\n%s\n\n", result)
	
	// Test non-existing check
	result = ExtractCheckInfo(sampleInput, "nonexistent")
	fmt.Printf("Result for 'nonexistent': '%s'\n\n", result)
	
	// Test first check
	result = ExtractCheckInfo(sampleInput, "process_agent")
	fmt.Printf("Result for 'process_agent':\n%s\n\n", result)
	
	// Test last check
	result = ExtractCheckInfo(sampleInput, "uptime")
	fmt.Printf("Result for 'uptime':\n%s\n", result)

	// Test vscode check
	result = ExtractCheckInfo(sampleInput, "vscode")
	fmt.Printf("Result for 'vscode':\n%s\n", result)
}
