package commands

import (
	"testing"
)

func TestShowEnv(t *testing.T) {
	cmd := ShowEnv()

	// Test case: Check if the command use is set correctly
	if cmd.Use != "showEnv" {
		t.Errorf("Expected command use to be 'showEnv', but got '%s'", cmd.Use)
	}

	// Test case: Check if the command short description is set correctly
	if cmd.Short != "Show environment variables and values" {
		t.Errorf("Expected command short description to be 'Show environment variables and values', but got '%s'", cmd.Short)
	}

	// Test case: Check if the command long description is set correctly
	expectedLong := "Show environment variables and values, show as key=value pairs."
	if cmd.Long != expectedLong {
		t.Errorf("Expected command long description to be '%s', but got '%s'", expectedLong, cmd.Long)
	}

	// Test case: Check if the command run function is set correctly
	// if cmd.Run != showEnv {
	// 	t.Errorf("Expected command run function to be 'showEnv', but got '%v'", cmd.Run)
	// }
}
