package main

import (
	"fmt"
	filesvalidator "headgen/internal/files-validator"
	"headgen/internal/headers"
	"os"
)

func main() {
	fileNames, err := filesvalidator.CheckFilesArgs(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, fileName := range fileNames {
		err = headers.Generate(fileName)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
