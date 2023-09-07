package commands

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"
)

func ShowEnv() *cobra.Command {
	return &cobra.Command{
		Use:   "showEnv",
		Short: "Show environment variables and values",
		Long:  `Show environment variables and values, show as key=value pairs.`,
		Run:   showEnv,
	}
}

func showEnv(cmd *cobra.Command, args []string) {
	file := getEnv()

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
