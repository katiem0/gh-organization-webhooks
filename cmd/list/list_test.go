package list

import (
	"testing"
)

func TestNewCmdList(t *testing.T) {
	cmd := NewCmdList()

	if cmd == nil {
		t.Fatal("NewCmdList() returned nil")
	}

	// Test basic properties
	if cmd.Use != "list <source organization> [flags]" {
		t.Errorf("Expected Use to be 'list <source organization> [flags]', got %s", cmd.Use)
	}

	// Test flags
	if cmd.Flag("hostname") == nil {
		t.Error("hostname flag not found")
	}

	if cmd.Flag("output-file") == nil {
		t.Error("output-file flag not found")
	}

	if cmd.Flag("token") == nil {
		t.Error("token flag not found")
	}

	// Test short description
	if cmd.Short == "" {
		t.Error("Command should have a short description")
	}
}
