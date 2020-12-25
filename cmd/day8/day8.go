package main

import (
	"advent2020/internal/util"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Hello from day 8")
	var program, err = util.GetLines("day8", "boot.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	var compiledProgram = make([]*instruction, 0, len(program))
	for _, line := range program {
		var instr, err = toInstruction(line)
		//fmt.Printf("\n%#v", instr)
		if err != nil {
			panic("Failed to parse line:" + line)
		}
		compiledProgram = append(compiledProgram, instr)
	}
	var state = newState()
	var finalState = run(*state, compiledProgram)
	fmt.Printf("\nAccumulator at the end is:%d", finalState.acc)

	for idx, compiledLine := range compiledProgram {
		var op = compiledLine.opp
		if op == nop || op == jmp {
			switchInstruction(compiledProgram, idx)
			var state = newState()
			finalState = run(*state, compiledProgram)
			switchInstruction(compiledProgram, idx)
			if finalState.atEnd {
				fmt.Printf("\nGot to the end with accumulatr:%d and switch:%d", finalState.acc, idx)
			}
		}
	}
}

type iState struct {
	acc   int
	lines map[int]bool
	atEnd bool
}

func switchInstruction(compileProgram []*instruction, idx int) {
	switch compileProgram[idx].opp {
	case nop:
		compileProgram[idx].opp = jmp
	case jmp:
		compileProgram[idx].opp = nop
	default:
		fmt.Printf("Switch called for invalid instruction idx:%d,  %#v", idx, compileProgram[idx])
	}
}

func newState() *iState {
	var retVal = iState{
		acc:   0,
		lines: make(map[int]bool),
		atEnd: false,
	}
	return &retVal
}

type op int

func (this op) String() string {
	switch this {
	case acc:
		return "acc"
	case nop:
		return "nop"
	case jmp:
		return "jmp"
	default:
		return "Invalid op" + string(this)
	}
}

const (
	acc = op(1)
	jmp = op(2)
	nop = op(3)
)

type instruction struct {
	param int
	opp   op
}

var instrRe, _ = regexp.Compile(`((nop)|(acc)|(jmp))\s([+-])(\d+)`)

func toInstruction(line string) (*instruction, error) {
	var matches = instrRe.FindStringSubmatch(line)
	if len(matches) != 7 {
		return nil, errors.New("Invalid instruction:" + line)
	}
	var opp op
	switch matches[1] {
	case "nop":
		opp = nop
	case "jmp":
		opp = jmp
	case "acc":
		opp = acc
	default:
		return nil, errors.New("Invalid op code:" + matches[1])
	}

	var sign int64
	switch matches[5] {

	case "+":
		sign = 1
	case "-":
		sign = -1
	default:
		return nil, errors.New("Invalid sign simbol:" + matches[5])
	}

	var param, err = strconv.ParseInt(matches[6], 10, 32)
	if err != nil {
		return nil, errors.New("Invalid param in op" + matches[6])
	}
	param = param * sign

	var retVal = instruction{
		param: int(param),
		opp:   opp,
	}
	return &retVal, nil

}

func run(state iState, program []*instruction) iState {
	var ip = 0

	for !state.lines[ip] {
		if ip < 0 || ip > len(program) {
			panic("Out of memory, trying to execute line" + string(ip))
		}
		if ip == len(program) {
			state.atEnd = true
			return state
		}
		state.lines[ip] = true
		var current = program[ip]
		switch current.opp {
		case nop:
			ip += 1
		case jmp:
			ip += current.param
		case acc:
			ip += 1
			state.acc += current.param
		}
	}
	return state
}
