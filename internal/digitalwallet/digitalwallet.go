// internal/digitalwallet/digitalwallet.go
package digitalwallet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

// App represents the main application
type App struct {
	Verbose       bool
	ProcessedCount int
}

// ProcessResult represents processing results
type ProcessResult struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

// NewApp creates a new application instance
func NewApp(verbose bool) *App {
	return &App{
		Verbose:       verbose,
		ProcessedCount: 0,
	}
}

// Run executes the main application logic
func (a *App) Run(inputFile, outputFile string) error {
	if a.Verbose {
		log.Println("Starting DigitalWallet processing...")
	}

	// Read input data
	var inputData string
	if inputFile != "" {
		if a.Verbose {
			log.Printf("Reading from file: %s", inputFile)
		}
		data, err := ioutil.ReadFile(inputFile)
		if err != nil {
			return fmt.Errorf("failed to read input file: %w", err)
		}
		inputData = string(data)
	} else {
		inputData = "Sample data for processing"
		if a.Verbose {
			log.Println("Using default test data")
		}
	}

	// Process the data
	result, err := a.Process(inputData)
	if err != nil {
		return fmt.Errorf("processing failed: %w", err)
	}

	// Generate output
	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal result: %w", err)
	}

	// Save or print output
	if outputFile != "" {
		if a.Verbose {
			log.Printf("Writing results to: %s", outputFile)
		}
		err = ioutil.WriteFile(outputFile, output, 0o644)
		if err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}
	} else {
		fmt.Println(string(output))
	}

	if a.Verbose {
		log.Printf("Processing complete. Total processed: %d", a.ProcessedCount)
	}

	return nil
}

// Process handles the core data processing
func (a *App) Process(data string) (*ProcessResult, error) {
	if a.Verbose {
		log.Printf("Processing data of length: %d", len(data))
	}

	// Simulate processing
	a.ProcessedCount++

	result := &ProcessResult{
		Success:   true,
		Message:   fmt.Sprintf("Successfully processed item #%d", a.ProcessedCount),
		Data: map[string]interface{}{
			"length":       len(data),
			"processed_at": time.Now().Format(time.RFC3339),
			"item_number":  a.ProcessedCount,
		},
		Timestamp: time.Now(),
	}

	return result, nil
}

// GetStats returns application statistics
func (a *App) GetStats() map[string]interface{} {
	return map[string]interface{}{
		"processed_count": a.ProcessedCount,
		"verbose":        a.Verbose,
	}
}
