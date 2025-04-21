package main

import (
	"fmt"
	filesvalidator "headgen/internal/files-validator"
	"os"
)

func main() {
	fileNames, err := filesvalidator.CheckFilesArgs(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Provided files: %s", fileNames)
}
