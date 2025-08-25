package cmd

import (
	"testing"
)

func TestNewCmd(t *testing.T) {
	cmd := NewCmd()

	if cmd == nil {
		t.Fatal("NewCmd() returned nil")
	}

	// Test basic properties
	if cmd.Use != "organization-webhooks <command> [flags]" {
		t.Errorf("Expected Use to be 'organization-webhooks <command> [flags]', got %s", cmd.Use)
	}

	// Test subcommands
	var hasListCmd, hasCreateCmd bool
	for _, subCmd := range cmd.Commands() {
		if subCmd.Name() == "list" {
			hasListCmd = true
		}
		if subCmd.Name() == "create" {
			hasCreateCmd = true
		}
	}

	if !hasListCmd {
		t.Error("Missing 'list' subcommand")
	}

	if !hasCreateCmd {
		t.Error("Missing 'create' subcommand")
	}

	// Test short description
	if cmd.Short == "" {
		t.Error("Command should have a short description")
	}

	// Test long description
	if cmd.Long == "" {
		t.Error("Command should have a long description")
	}

	// Test completion options
	if !cmd.CompletionOptions.DisableDefaultCmd {
		t.Error("Default completion command should be disabled")
	}
}
