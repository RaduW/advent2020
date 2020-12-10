package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func GetLines(fileName string) ([]string, error) {
	var absPath, err = filepath.Abs(fileName)
	if err != nil {
		return nil, err
	}
	var file *os.File
	file, err = os.Open(absPath)
	defer file.Close()

	if err != nil {
		fmt.Println("\nCould not open file:", absPath)
		return nil, err
	}

	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var retVal = make([]string, 0, 100)
	for scanner.Scan() {
		retVal = append(retVal, scanner.Text())
	}

	return retVal, nil
}
