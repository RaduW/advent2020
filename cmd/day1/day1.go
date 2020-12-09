package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	var x string
	var err error
	x, err = os.Executable()
	if err == nil {
		fmt.Printf("The executable is %v.\n", x)
	} else {
		fmt.Printf("we have an error %v.\n", err)
	}
	x, err = os.Getwd()
	if err == nil {
		fmt.Printf("The working directory is %v.\n", x)
	} else {
		fmt.Printf("we have an error %v.\n", err)
	}
	fmt.Printf("hello")

	var expenses = getExpenses()
	var v1, v2 = findExpensePair(expenses)
	fmt.Printf("v1=%d v2=%d v1+v2=%d v1*v2=%d\n", v1, v2, v1+v2, v1*v2)
	var v3 int64
	v1, v2, v3 = findExpenseTriple(expenses)
	fmt.Printf("v1=%d v2=%d v3, v1+v2+v3=%d v1*v2*v3=%d\n", v1, v2, v1+v2+v3, v1*v2*v3)
}

func printExpenses(expenses []int64) {
	for _, expense := range expenses {
		fmt.Printf("->%d<-\n", expense)
	}
}

func getLines(fileName string) error {
	var absPath, err = filepath.Abs(fileName)
	if err != nil {
		return err
	}
	var file *os.File
	file, err = os.Open(absPath)
	defer file.Close()

	if err != nil {
		fmt.Println("\nCould not open file:", absPath)
		return err
	}

	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}

func findExpensePair(expenses []int64) (int64, int64) {
	for idx, val1 := range expenses {
		var left = 2020 - val1
		if left > 0 {
			for _, val2 := range expenses[idx+1:] {
				if val2 == left {
					return val1, val2
				}
			}

		}
	}
	return -1, -1
}
func findExpenseTriple(expenses []int64) (int64, int64, int64) {
	for idx1, val1 := range expenses {
		var left1 = 2020 - val1
		if left1 > 0 {
			for idx2, val2 := range expenses[idx1+1:] {
				var left2 = left1 - val2
				if left2 > 0 {
					for _, val3 := range expenses[idx2+1:] {
						if val3 == left2 {
							return val1, val2, val3
						}
					}
				}
			}
		}
	}
	return -1, -1, -1
}

func getExpenses() (retVal []int64) {
	retVal = make([]int64, 0, 100)
	var absPath, err = filepath.Abs("expenses.txt")
	if err != nil {
		return
	}
	var file *os.File
	file, err = os.Open(absPath)
	defer file.Close()

	if err != nil {
		return
	}
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var val int64
		val, _ = strconv.ParseInt(scanner.Text(), 10, 64)
		if val != 0 {
			retVal = append(retVal, val)
		}
	}
	return
}
