package commands

import (
	"testing"
)

func TestGenerateKey(t *testing.T) {
	// Test case: Verify the generated command has the correct name
	t.Run("CommandName", func(t *testing.T) {
		cmd := GenerateKey()
		expectedName := "generateKey"
		if cmd.Use != expectedName {
			t.Errorf("Expected command name to be %s, but got %s", expectedName, cmd.Use)
		}
	})

	// Test case: Verify the generated command has the correct short description
	t.Run("CommandShortDescription", func(t *testing.T) {
		cmd := GenerateKey()
		expectedShort := "Generate a new app key"
		if cmd.Short != expectedShort {
			t.Errorf("Expected command short description to be %s, but got %s", expectedShort, cmd.Short)
		}
	})

	// Test case: Verify the generated command has the correct long description
	t.Run("CommandLongDescription", func(t *testing.T) {
		cmd := GenerateKey()
		expectedLong := "Look in env file where APP_KEY is located and add new value"
		if cmd.Long != expectedLong {
			t.Errorf("Expected command long description to be %s, but got %s", expectedLong, cmd.Long)
		}
	})

	// Test case 4: Verify the command's Run field
	// t.Run("CommandRun", func(t *testing.T) {
	// 	cmd := GenerateKey()
	// 	expected := generateKey
	// })
}
