package filesvalidator

import (
	"fmt"
	"os"
	"path/filepath"
)

func CheckFilesArgs(args []string) ([]string, error) {
	if len(args) <= 1 {
		return nil, fmt.Errorf("usage: headgen file1.c file2.c ...")
	}

	fileNames := args[1:]
	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			return nil, fmt.Errorf("no such file: %s", fileName)
		} else if filepath.Ext(file.Name()) != ".c" {
			file.Close()
			return nil, fmt.Errorf("only .c files supported, you provided: %s", fileName)
		}
		file.Close()
	}

	return fileNames, nil
}
