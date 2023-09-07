package main

import (
	"github.com/andresmeireles/resume/internal/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(
		commands.ShowEnv(),
		commands.GenerateKey(),
		commands.ExecuteMigrations(),
		commands.RollbackMigration(),
		commands.CreateUser(),
	)
	rootCmd.Execute()
}
