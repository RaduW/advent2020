package main

import (
	"advent2020/internal/util"
	"fmt"
	"strconv"
)

func main() {
	var expenses = getExpenses()
	var v1, v2 = findExpensePair(expenses)
	fmt.Printf("v1=%d v2=%d v1+v2=%d v1*v2=%d\n", v1, v2, v1+v2, v1*v2)
	var v3 int64
	v1, v2, v3 = findExpenseTriple(expenses)
	fmt.Printf("v1=%d v2=%d v3, v1+v2+v3=%d v1*v2*v3=%d\n", v1, v2, v1+v2+v3, v1*v2*v3)
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

	var lines, err = util.GetLines("expenses.txt")

	if err != nil {
		return
	}
	for _, line := range lines {
		var val int64
		val, _ = strconv.ParseInt(line, 10, 64)
		if val != 0 {
			retVal = append(retVal, val)
		}

	}

	return
}
