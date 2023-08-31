package commands

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func GenerateKey() *cobra.Command {
	return &cobra.Command{
		Use:   "generateKey",
		Short: "Generate a new app key",
		Long:  `Look in env file where APP_KEY is located and add new value`,
		Run:   generateKey,
	}
}

func generateKey(cmd *cobra.Command, args []string) {
	file := getEnv()

	tempFile, err := os.CreateTemp("", "key")

	if err != nil {
		panic(err)
	}

	defer file.Close()
	defer tempFile.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		hasAppKey := strings.Contains(line, "APP_KEY=")
		if hasAppKey {
			fmt.Fprintln(tempFile, "APP_KEY="+keyGenration())
			continue
		}
		fmt.Fprintln(tempFile, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if err := file.Close(); err != nil {
		panic(err)
	}

	if err := os.Remove(file.Name()); err != nil {
		panic(err)
	}

	if err := os.Rename(tempFile.Name(), file.Name()); err != nil {
		panic(err)
	}
}

func keyGenration() string {
	str := time.Now().String()

	return base64.StdEncoding.EncodeToString([]byte(str))
}
