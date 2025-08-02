// internal/digitalwallet/digitalwallet_test.go
package digitalwallet

import (
	"testing"
)

func TestNewApp(t *testing.T) {
	app := NewApp(true)
	if app == nil {
		t.Fatal("NewApp returned nil")
	}
	if !app.Verbose {
		t.Error("Expected verbose to be true")
	}
	if app.ProcessedCount != 0 {
		t.Errorf("Expected ProcessedCount to be 0, got %d", app.ProcessedCount)
	}
}

func TestProcess(t *testing.T) {
	app := NewApp(false)
	result, err := app.Process("test data")
	
	if err != nil {
		t.Fatalf("Process returned error: %v", err)
	}
	
	if !result.Success {
		t.Error("Expected result.Success to be true")
	}
	
	if app.ProcessedCount != 1 {
		t.Errorf("Expected ProcessedCount to be 1, got %d", app.ProcessedCount)
	}
}

func TestRun(t *testing.T) {
	app := NewApp(false)
	err := app.Run("", "")
	
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
}
