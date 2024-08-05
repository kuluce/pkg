package logger

import (
	"fmt"
	"testing"
)

func TestInitLog(t *testing.T) {
	// Call the InitLog function
	InitLog("test")

	// Assert that the logger variable is not nil
	if logger == nil {
		t.Errorf("Expected logger to be initialized, but got nil")
	}

	// Add more assertions if needed
}

func SetUp() {
	InitLog("test")
}

func TestInfo(t *testing.T) {
	SetUp()
	fmt.Print("TestInfo\n")
	Info("test")
}
