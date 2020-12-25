package main

import (
	"advent2020/internal/util"
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("hello day 7 ")
	var lines, _ = util.GetLines("day7", "bagRules.txt")
	var allBags, _ = getAllBags(lines)
	for col, bag := range allBags {
		fmt.Printf("\n%v=>%v", col, bag)
	}

	var c = color("shiny gold")
	fmt.Printf("\n\nParents for %v:\n", c)
	var parents, err = getAllParents(c, allBags)
	if err != nil {
		fmt.Println("Error while looking for parents")
	}
	fmt.Printf("\n!!!! Found %d parents:\n", len(parents))
	for _, bag := range parents {
		fmt.Println(bag)
	}
	fmt.Printf("\n!!!! Found %d parents from a total of %d bag colors.\n", len(parents), len(allBags))
	var childCount int
	//var c2 = c
	//c2 = color("dark olive")
	childCount, err = getChildrenCount(c, allBags)
	fmt.Printf("\n!!!! Found %d children for bag color %s\n", childCount, c)

}

type bagSet map[color]*bagRelation

func (this bagSet) getBag(cl color) *bagRelation {
	var retVal = this[cl]
	if retVal == nil {
		retVal = new(bagRelation)
		this[cl] = retVal
		retVal.color = cl
		retVal.parents = make([]*bagRelation, 0)
		retVal.children = make([]childRelation, 0)
	}
	return retVal
}

type color string

type bagRelation struct {
	color    color
	parents  []*bagRelation
	children []childRelation
}

func (this bagRelation) String() string {
	var buffer bytes.Buffer

	fmt.Fprintf(&buffer, "%v\n", this.color)
	fmt.Fprintf(&buffer, "   children: [ ")
	for idx, child := range this.children {
		if idx != 0 {
			fmt.Fprintf(&buffer, ", ")
		}
		fmt.Fprintf(&buffer, "%d->%s", child.count, child.child.color)
	}
	fmt.Fprintf(&buffer, "]\n")
	fmt.Fprintf(&buffer, "   parents: [ ")
	for idx, parent := range this.parents {
		if idx != 0 {
			fmt.Fprintf(&buffer, ", ")
		}
		fmt.Fprintf(&buffer, "%s", parent.color)
	}
	fmt.Fprintf(&buffer, "]")
	return buffer.String()

}

type childRelation struct {
	count int
	child *bagRelation
}

type bagSpec struct {
	color    color
	children []childSpec
}

type childSpec struct {
	color color
	count int
}

func getAllBags(lines []string) (bagSet, error) {
	var ab = make(bagSet)
	for _, line := range lines {
		var bs, err = readBagSpec(line)
		if err != nil {
			return nil, err
		}
		var currentBag = ab.getBag(bs.color)
		for _, cs := range bs.children {
			var child = ab.getBag(cs.color)
			currentBag.children = append(currentBag.children, childRelation{child: child, count: cs.count})
			child.parents = append(child.parents, currentBag)
		}
	}

	return ab, nil
}

func readBagSpec(line string) (*bagSpec, error) {
	var parsedSpec = bagSpecRe.FindStringSubmatch(line)
	if len(parsedSpec) != 5 {
		return nil, errors.New("invalid spec:" + line)
	}
	var cl = parsedSpec[1]
	var retVal = new(bagSpec)
	retVal.color = color(cl)
	retVal.children = make([]childSpec, 0)
	if len(parsedSpec[4]) != 0 {
		// parse children
		var chs = childrenSpecRe.FindAllStringSubmatch(parsedSpec[4], -1)
		for _, cs := range chs {
			if len(cs) != 4 {
				return nil, errors.New("invalid child spec" + parsedSpec[4])
			}
			var count, _ = strconv.ParseInt(cs[1], 10, 64)
			retVal.children = append(retVal.children, childSpec{color: color(cs[2]), count: int(count)})
		}
	}

	return retVal, nil

}

func getChildrenCount(c color, allBags bagSet) (int, error) {
	var retVal int

	var current = allBags[c]
	if current == nil {
		return 0, errors.New("Unknown bag of color:" + string(c))
	}
	for _, cr := range current.children {
		var numChildren = cr.count
		var cnt, err = getChildrenCount(cr.child.color, allBags)

		if err != nil {
			return 0, err
		}

		retVal += (cnt + 1) * numChildren
	}

	return retVal, nil
}

func getAllParents(c color, allBags bagSet) ([]*bagRelation, error) {
	var parents = make(bagSet)

	var getAllParentsInternal func(child *bagRelation) error

	getAllParentsInternal = func(child *bagRelation) error {
		if child == nil {
			return errors.New("could not find original color in all bags")
		}

		for _, parent := range child.parents {
			var pColor = parent.color
			if parents[pColor] != nil {
				//already taken care of this parent
				continue
			}
			parents[pColor] = parent
			var err = getAllParentsInternal(parent)
			if err != nil {
				return err
			}
		}

		if parents[c] == nil {
			//already taken care of this color nothing left to do
			return nil
		}

		return nil
	}

	var child = allBags[c]
	var err = getAllParentsInternal(child)

	if err != nil {
		return nil, err
	}

	var retVal = make([]*bagRelation, 0, len(parents))
	for _, bag := range parents {
		retVal = append(retVal, bag)
	}
	return retVal, nil
}

var bagSpecRe, _ = regexp.Compile(`^([a-z]+\s+[a-z]+) bags contain ((no other bags.)|(\d+.+))$`)
var childrenSpecRe, _ = regexp.Compile(`(\d+) ([a-z]+\s+[a-z]+) bags?(, |\.)`)
