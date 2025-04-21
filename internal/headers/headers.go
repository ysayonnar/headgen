package headers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// #ifndef FUNCTIONS_H
// #define FUNCTIONS_H
//

func Generate(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error while opening file: %w", err)
	}
	defer file.Close()

	clearName := strings.TrimSuffix(filename, filepath.Ext(filename))

	headerFileName := clearName + ".h"
	headerFile, err := os.Create(headerFileName)
	if err != nil {
		return fmt.Errorf("error while creating header file: %w", err)
	}
	defer headerFile.Close()

	//adding head of header file
	libName := strings.ToUpper(clearName) + "_H"
	fmt.Fprintf(headerFile, "#ifndef %s\n", libName)
	fmt.Fprintf(headerFile, "#define %s\n", libName)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		headerLine, err := parseLine(line)
		if err != nil {
			continue
		}
		_, err = fmt.Fprintln(headerFile, headerLine)
		if err != nil {
			return fmt.Errorf("error while writing to the .h file: %w", err)
		}
	}

	if err = scanner.Err(); err != nil {
		return fmt.Errorf("error while reading source file: %w", err)
	}

	fmt.Fprintln(headerFile, "#endif")

	return nil
}

func parseLine(line string) (string, error) {
	if strings.HasPrefix(line, "#include") || strings.HasPrefix(line, "#define") {
		return line, nil
	}

	//регулярное выражение - по другому не придумал как парсить
	pattern := `^\s*\w+\s+\w+\s*\([\w\s,*\[\].=&]*\)\s*\{\s*$`
	matched, err := regexp.MatchString(pattern, line)
	if err != nil || !matched {
		return "", errors.New("dont matches")
	} else {
		return strings.Replace(line, "{", ";", 1), nil
	}
}
