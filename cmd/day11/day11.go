package main

import (
	"advent2020/internal/util"
	"errors"
	"fmt"
)

type placeType rune

const (
	reserved placeType = '.'
	free     placeType = 'L'
	occupied placeType = '#'
)

type seatMap [][]placeType

func main() {
	fmt.Println("Hello day 11")

	var lines, _ = util.GetLines("day11", "seatMap.txt")
	var source, err = toSeatMap(lines)

	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}

	var mp2 seatMap

	for _, evolveAlgo := range [](func(seatMap, int, int) placeType){evolveCell1, evolveCell2} {
		var mp1 = make(seatMap, len(source))
		for idx, row := range source {
			mp1[idx] = make([]placeType, len(source[idx]))
			copy(mp1[idx], row)
		}
		resetSeatMap(mp2)
		var count = 0
		for areDifferent(mp1, mp2) {
			mp2, err = evolve(mp1, mp2, evolveAlgo)
			if err != nil {
				fmt.Println(err)
				return
			}
			count++
			mp1, mp2 = mp2, mp1

		}
		fmt.Printf("\nNumber of iterations: %d", count)
		var occ = countSeats(mp1, occupied)
		fmt.Printf("\nNumber of occupied seats: %d", occ)
		var f = countSeats(mp1, free)
		fmt.Printf("\nNumber of free seats: %d", f)
		var r = countSeats(mp1, reserved)
		fmt.Printf("\nNumber of reserved seats: %d", r)
		fmt.Printf("\nTotal: %d+%d+%d=%d %d*%d=%d\n\n", occ, f, r, occ+f+r, len(mp1), len(mp1[0]), len(mp1)*len(mp1[0]))

	}
}

func resetSeatMap(m seatMap) {
	if m == nil {
		return
	}
	for _, line := range m {
		for idx, _ := range line {
			line[idx] = free
		}
	}
}

func countSeats(m seatMap, state placeType) int {
	var retVal = 0
	for _, line := range m {
		for _, ch := range line {
			if ch == state {
				retVal++
			}
		}
	}
	return retVal
}

func display(m seatMap, count int) {
	if len(m) > 0 {
		for _, _ = range m[0] {
			fmt.Print("--")
		}
	}
	fmt.Printf("\n\n\t%d\n", count)
	for _, line := range m {
		for _, ch := range line {
			fmt.Printf("%c ", ch)
		}
		fmt.Printf("\n")
	}
}

func areDifferent(mp1, mp2 seatMap) bool {
	if mp1 == nil || mp2 == nil {
		return true
	}

	if len(mp1) != len(mp2) {
		return true
	}

	for rowIdx, line := range mp1 {
		if len(line) != len(mp2[rowIdx]) {
			return true
		}
		for colIdx, elm := range line {
			if elm != mp2[rowIdx][colIdx] {
				return true
			}
		}
	}
	return false
}

func evolve(currentM seatMap, newM seatMap, evolveCell func(seatMap, int, int) placeType) (seatMap, error) {
	if currentM == nil {
		return nil, errors.New("nil map passed into evolve")
	}
	if newM == nil {
		newM = make(seatMap, len(currentM), len(currentM))
		for idx, row := range currentM {
			(newM)[idx] = make([]placeType, len(row), len(row))
		}
	} else {
		//check sizes for compatibility
		if len(currentM) != len(newM) {
			return nil, errors.New("old and new map have different sizes")
		}
		for idx, row := range currentM {
			if len(newM[idx]) != len(row) {
				return nil, errors.New(fmt.Sprintf("row %d has different sizes in old and new", idx))
			}
		}
	}
	//also check for square maps ( could be done in one go with the checks above but this is clearer)
	var length = -1
	for idx, row := range currentM {
		if idx == 0 {
			length = len(row)
		} else {
			if length != len(row) {
				return nil, errors.New("non rectangular map received")
			}
		}
	}

	for rowIdx, row := range currentM {
		for colIdx, _ := range row {
			newM[rowIdx][colIdx] = evolveCell(currentM, rowIdx, colIdx)
		}
	}

	return newM, nil
}

func evolveCell2(m seatMap, r, c int) placeType {
	var deltas = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	var maxR = len(m) - 1
	var maxC = len(m[0]) - 1

	if m[r][c] == reserved {
		return reserved
	}
	var neighboursCount = 0

	for _, delta := range deltas {
		var rd = r + delta[0]
		var cd = c + delta[1]
	stopLooping:
		for {
			if rd < 0 || rd > maxR || cd < 0 || cd > maxC {
				break stopLooping
			}

			switch m[rd][cd] {
			case reserved:
				rd += delta[0]
				cd += delta[1]
			case occupied:
				neighboursCount++
				break stopLooping
			case free:
				break stopLooping
			default:
				fmt.Printf("Invalid %c", m[rd][cd])
			}
		}
	}
	if m[r][c] == free && neighboursCount == 0 {
		return occupied
	}
	if m[r][c] == occupied && neighboursCount >= 5 {
		return free
	}
	return m[r][c]

}

func evolveCell1(m seatMap, r, c int) placeType {
	var deltas = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	var maxR = len(m) - 1
	var maxC = len(m[0]) - 1

	if m[r][c] == reserved {
		return reserved
	}

	var neighboursCount = 0
	for _, delta := range deltas {
		var rd = r + delta[0]
		var cd = c + delta[1]

		if rd < 0 || rd > maxR || cd < 0 || cd > maxC {
			continue
		}
		if m[rd][cd] == occupied {
			neighboursCount++
		}
	}

	if m[r][c] == free && neighboursCount == 0 {
		return occupied
	}
	if m[r][c] == occupied && neighboursCount >= 4 {
		return free
	}
	return m[r][c]
}

func toSeatMap(lines []string) (seatMap, error) {
	var retVal seatMap = make(seatMap, 0, len(lines))
	for lineIdx, line := range lines {
		for charIdx, chr := range line {
			switch placeType(chr) {
			case free, occupied, reserved:
				continue
			default:
				return nil, errors.New(fmt.Sprintf("Invalid map entry at line:%d column:%d %c", lineIdx, charIdx, chr))
			}
		}
		retVal = append(retVal, ([]placeType)(line))
	}
	return retVal, nil
}
