package main

import (
	"advent2020/internal/util"
	"errors"
	"fmt"
	"regexp"
	"sort"
)

func main() {
	fmt.Println("hello day 5 ")
	var lines, _ = util.GetLines("boardingTickets.txt")
	var maxSeatId = 0
	var rowCols = make([]*rowCol, 0, len(lines))
	for _, line := range lines {
		var rowCol, _ = str2rowCol(line)
		var seatId = rowCol.seatId()
		if seatId > maxSeatId {
			maxSeatId = seatId
		}
		fmt.Printf("\nboardPass=%s row=%d col=%d seatId=%d", line, rowCol.row, rowCol.col, seatId)
		rowCols = append(rowCols, rowCol)
	}

	sort.Sort(bySeatId(rowCols))
	var emptySeatId int
	for idx, rc := range rowCols {
		if idx > 0 {
			if rowCols[idx-1].seatId()+1 != rowCols[idx].seatId() {
				emptySeatId = rowCols[idx].seatId() - 1
				fmt.Printf("\n\n EMPTY SEAT %d\n\n", emptySeatId)
			}

		}
		fmt.Printf("\nrow:%d col:%d setId:%d", rc.row, rc.col, rc.seatId())
	}

	fmt.Printf("\n Max seat id is %d", maxSeatId)
	fmt.Printf("\n Empty seat id is %d", emptySeatId)

}

type rowCol struct {
	row int
	col int
}

type bySeatId []*rowCol

func (a bySeatId) Len() int           { return len(a) }
func (a bySeatId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySeatId) Less(i, j int) bool { return a[i].seatId() < a[j].seatId() }

func (this *rowCol) seatId() int {
	return this.row*8 + this.col
}

var boardPassRe, _ = regexp.Compile(`^((F|B){7})((R|L){3})$`)

func str2rowCol(str string) (*rowCol, error) {
	var subMatches = boardPassRe.FindStringSubmatch(str)
	if len(subMatches) != 5 {
		return nil, errors.New("invalid boarding pass:" + str)
	}
	var rowStr, colStr = subMatches[1], subMatches[3]
	var row = 0
	for _, ch := range rowStr {
		row = row * 2
		if ch == 'B' {
			row += 1
		}
	}
	var col = 0
	for _, ch := range colStr {
		col = col * 2
		if ch == 'R' {
			col += 1
		}
	}
	return &rowCol{row, col}, nil
}
