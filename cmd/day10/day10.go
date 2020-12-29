package main

import (
	"advent2020/internal/util"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Hello from day 10")

	var fileName = "test2Jolts.txt"
	var day = "day10"
	var lines, _ = util.GetLines(day, fileName)
	var vals = make([]int, 0, len(lines)+1)
	vals = append(vals, 0)
	for _, line := range lines {
		var val, err = strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Printf("\nError parsing line:\n->%s<-\n", line)
		}
		vals = append(vals, int(val))
	}
	sort.Ints(vals)
	var sortedFname = "sorted_" + fileName
	var path = filepath.Join("cmd", day, sortedFname)

	var f, err = os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, val := range vals {
		fmt.Fprintf(f, "%d\n", val)
	}
	var distCounts = make(map[int]int)
	for idx := 0; idx < len(vals); idx++ {
		var diff int
		if idx == 0 {
			diff = vals[0]
		} else {
			diff = vals[idx] - vals[idx-1]
		}
		distCounts[diff] += 1
	}
	// add the builtin addapter
	distCounts[3] += 1

	for diff, count := range distCounts {
		fmt.Printf("\n %d adapters have a diff of %d", count, diff)
	}

	fmt.Printf("\n")
	for _, val := range vals {
		fmt.Printf(" %d", val)
	}
	var seqGroups = countSequenceGroups(vals)
	fmt.Printf("\nSequenceGroups:%#v", seqGroups)
	var combinations = 1
	for length, count := range seqGroups {
		combinations *= intPow(count1SequenceVariations(length), count)
	}
	fmt.Printf("\nNumber of combinations:%d", countAdapterCombos(vals))
}
func intPow(num int, expo int) int {
	var retVal = num
	for i := 2; i <= expo; i++ {
		retVal *= num
	}
	return retVal
}

func countAdapterCombos(adapters []int) int {
	var seqGroups = countSequenceGroups(adapters)
	fmt.Printf("\nSequenceGroups:%#v", seqGroups)
	var combinations = 1
	for length, count := range seqGroups {
		combinations *= intPow(count1SequenceVariations(length), count)
	}
	return combinations
}

// Returns a dictionary of number of sequences of n consecutive numbers
// i.e. { 3:5 , 4:7} --> 5 sequences of 3 and 7 sequences of 4
func countSequenceGroups(adapters []int) map[int]int {
	var retVal = map[int]int{}
	var seqLen = 0
	for i := 1; i < len(adapters)-1; i++ {
		var x = adapters[i]
		var prev = adapters[i-1]
		var next = adapters[i+1]
		if x-prev == 1 && next-x == 1 {
			seqLen += 1
		} else {
			if seqLen > 0 {
				retVal[seqLen]++
				fmt.Printf("\nSequence of lenght %d finishing at %d", seqLen, prev)
			}
			seqLen = 0
		}
	}
	if seqLen > 0 {
		retVal[seqLen]++
	}

	return retVal
}

func count1SequenceVariations(numDigit int) int {
	switch numDigit {
	case 1:
		return 2
	case 2:
		return 4
	case 3:
		return 7
	}
	var n3 = 2
	var n2 = 4
	var n1 = 7
	var n int
	for i := 4; i <= numDigit; i++ {
		n = n1 + n2 + n3
		n3 = n2
		n2 = n1
		n1 = n
	}
	return n
}

func count1SequenceVariationsBruteForce(numDigits int) int {
	var max int64 = 1
	for i := 0; i < numDigits; i++ {
		max *= 2
	}
	max = max
	var v int64
	var retVal = 0
	for v = 0; v < max; v++ {
		var cnt = 0
		var res = v
		for div := 0; div < numDigits; div++ {
			var i = res % 2
			res = res / 2
			if i == 0 {
				cnt++
			} else {
				cnt = 0
			}
			if cnt == 3 {
				break
			}
		}
		if cnt < 3 {
			retVal++
		}
	}
	return retVal
}
