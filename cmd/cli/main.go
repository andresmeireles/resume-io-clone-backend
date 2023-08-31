package main

import (
	"github.com/andresmeireles/resume/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(commands.ShowEnv(), commands.GenerateKey())
	rootCmd.Execute()
}
