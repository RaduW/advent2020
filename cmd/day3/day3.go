package main

import (
	"advent2020/internal/util"
	"fmt"
)

func main() {
	fmt.Println("hello day 3 ")
	var lines, err = util.GetLines("day3", "treeMap.txt")
	if err != nil {
		fmt.Println("An error occurred while reading treeMap.txt", err)
	}

	for idx, line := range lines {
		fmt.Println(idx, line)
	}
	fmt.Println("")
	var hits = countHits(lines, 1, 2)
	fmt.Printf("\nNumber of hits = %d\n", hits)

	var speeds = [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	var all = 1
	for _, hv := range speeds {
		var horizontal, vertical = hv[0], hv[1]
		var hits = countHits(lines, horizontal, vertical)
		fmt.Printf("h:%d v:%d  hits=%d \n", horizontal, vertical, hits)
		all = all * hits
	}
	fmt.Printf("All hits: %d", all)
	//var hits1 = countHits(lines, 1,1)
	//var hits2 = countHits(lines, 1,1)

}

func countHits(lines []string, horizontalSpeed int, verticalSpeed int) (count int) {
	count = 0
	step := 0
	for idx := 0; idx < len(lines); idx += verticalSpeed {
		var line = lines[idx]
		if isHit(step, horizontalSpeed, line) {
			count++
		}
		step += 1
	}
	return
}

func isHit(step int, horizontalSpeed int, line string) bool {
	var space = '.'
	var tree = '#'
	var position = step * horizontalSpeed % len(line)

	var elm = int32(line[position])

	if elm == tree {
		return true
	} else {
		if elm != space {
			fmt.Println("MAP ERROR !!!!")
		}
		return false
	}

}
