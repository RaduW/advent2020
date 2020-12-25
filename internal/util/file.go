package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func GetLines(cmd string, fileName string) ([]string, error) {
	fileName = fmt.Sprintf("cmd/%s/%s", cmd, fileName)
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
