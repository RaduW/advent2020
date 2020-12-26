package main

import (
	"advent2020/internal/util"
	"errors"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello from day9")
	var lines, _ = util.GetLines("day9", "xmas.txt")
	var codes = make([]int64, len(lines))
	for idx, line := range lines {
		var val, err = strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Printf("\nError parsing line:\n->%s<-\n", line)
		}
		codes[idx] = val
	}

	var coder = newCoder(25)
	var badVal int64 = -1
	for _, val := range codes {
		var _, err = coder.addVal(val)
		badVal = val
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	var startIdx, endIdx, err = findSequenceIndexes(badVal, codes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nThe sequence adding up to %d is between [%d..%d] and min+max=%d", badVal, startIdx, endIdx, codes[startIdx]+codes[endIdx])
	fmt.Print("\n[")
	var smallest, largest int64
	smallest = codes[startIdx]
	largest = codes[startIdx]
	for idx := startIdx; idx <= endIdx; idx++ {
		var val = codes[idx]
		fmt.Printf("  %d", codes[idx])
		if val < smallest {
			smallest = val
		}
		if val > largest {
			largest = val
		}
	}
	fmt.Printf("]\n smallest + largest = %d + %d = %d", smallest, largest, smallest+largest)

}

type coder struct {
	buffer []int64
	index  int
	size   int
}

func (this *coder) addVal(val int64) (bool, error) {
	defer func() { this.index = (this.index + 1) % this.size }()
	if len(this.buffer) < this.size {
		this.buffer = append(this.buffer, val)
		return true, nil
	} else {
		for i1 := 0; i1 < len(this.buffer); i1++ {
			var v1 = this.buffer[i1]
			var needed = val - v1
			for i2 := i1 + 1; i2 < len(this.buffer); i2++ {
				if this.buffer[i2] == needed {
					this.buffer[this.index] = val
					return true, nil
				}
			}
		}
	}
	this.buffer[this.index] = val
	return false, errors.New("could not add val: " + strconv.Itoa(int(val)))
}

func newCoder(size int) *coder {
	var retVal = coder{
		buffer: make([]int64, 0, size),
		size:   size,
		index:  0,
	}
	return &retVal
}

func findSequenceIndexes(result int64, vals []int64) (int, int, error) {
	for startIdx, start := range vals {
		var sum = start
		var endIdx = startIdx + 1
		for sum < result {
			sum += vals[endIdx]
			if sum == result {
				// found result we can stop now
				return startIdx, endIdx, nil
			}
			endIdx += 1
		}
	}
	return 0, 0, errors.New(fmt.Sprintf("Could not find sequence for:%d ", result))

}
