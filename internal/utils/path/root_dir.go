package path

import (
	"os"
	"path/filepath"
)

func GetRootDir(cwd ...string) string {
	fileHasFolder := false
	currentDir := ""
	if len(cwd) > 0 {
		currentDir = cwd[0]
	}
	if currentDir == "" {
		cd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		currentDir = cd
	}

	err := filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == "go.mod" {
			fileHasFolder = true
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	if fileHasFolder {
		return currentDir
	}

	return GetRootDir(currentDir + "/..")
}
