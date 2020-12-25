package main

import (
	"advent2020/internal/util"
	"errors"
	"fmt"
)

type answerGroup struct {
	answers     asAnswers
	memberCount int
}

func main() {
	var lines, _ = util.GetLines("day6", "questions.txt")

	var answerGroups = toAnswerGroups(lines)
	var uniqueAnswersPerGroup = 0
	var uniqueCommonAnswersPerGroup = 0
	for _, ag := range answerGroups {
		var count, _ = ag.answers.distinctAnswersWithCount(ag.memberCount)
		fmt.Printf("\n%v, uniqueCount=%d", ag, count)
		var da, _ = ag.answers.distinctAnswers()
		var uca, _ = ag.answers.distinctAnswersWithCount(ag.memberCount)
		uniqueAnswersPerGroup += da
		uniqueCommonAnswersPerGroup += uca
	}
	fmt.Printf("\nNumber of groups=%d \nUnique answers per group=%d \nUnique common answers per group=%d ", len(answerGroups), uniqueAnswersPerGroup, uniqueCommonAnswersPerGroup)
}

func toAnswerGroups(lines []string) []*answerGroup {
	var retVal = make([]*answerGroup, 0, 100)
	var curAnswerGroup *answerGroup

	for _, line := range lines {
		if len(line) == 0 {
			curAnswerGroup = nil
		} else {
			if curAnswerGroup == nil {
				curAnswerGroup = new(answerGroup)
				curAnswerGroup.answers = make(map[rune]int)
				retVal = append(retVal, curAnswerGroup)
			}
			curAnswerGroup.memberCount += 1
			for _, r := range line {
				curAnswerGroup.answers.add(r)
			}
		}
	}
	return retVal
}

type answers interface {
	add(rune) error
	distinctAnswers() int
}

type asAnswers map[rune]int

func (c asAnswers) add(r rune) (int, error) {
	if c == nil {
		return 0, errors.New("invalid container")
	}
	var v = c[r] + 1
	c[r] = v

	return v, nil
}

func (c asAnswers) distinctAnswers() (int, error) {
	if c == nil {
		return 0, errors.New("invalid container")
	}
	return len(c), nil
}

func (c asAnswers) distinctAnswersWithCount(count int) (int, error) {
	if c == nil {
		return 0, errors.New("invalid container")
	}

	var retVal = 0
	for _, val := range c {
		if val == count {
			retVal += 1
		}
	}
	return retVal, nil
}
