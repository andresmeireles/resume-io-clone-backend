package commands

import "testing"

func TestGetEnv(t *testing.T) {
	// Test when .env file exists
	t.Run("GetEnv", func(t *testing.T) {
		cmd := getEnv()
		if cmd == nil {
			t.Error("Expected command to be not nil")
		}
	})
}
