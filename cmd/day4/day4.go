package main

import (
	"advent2020/internal/util"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello day 4 ")
	var re, _ = getPropertyRegex()
	var rawProps = getLineProperties(re, "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd")
	fmt.Println(rawProps)
	rawProps = getLineProperties(re, "  fasdf:a   ")
	fmt.Println("->", rawProps, "<-")
	var lines, _ = util.GetLines("day4", "passports.txt")
	var passports = parsePassports(lines)
	var numPass, validPass, invalidPass, superValid, superInvalid int
	numPass = len(passports)
	for _, passInfo := range passports {
		var valid string
		var isValid, errorStr = passInfo.isValid()

		if isValid {
			validPass += 1
			valid = "Valid passport: "
		} else {
			invalidPass += 1
			valid = "!!! Passport missing " + errorStr
		}
		//fmt.Printf("\n%s\n %+v\n", valid, passInfo)

		isValid, errorStr = passInfo.isSuperValid()

		if isValid {
			superValid += 1
			valid = "Super Valid passport: "
		} else {
			superInvalid += 1
			valid = "!!! Passport failing " + errorStr
		}
		fmt.Printf("\n%s\n %+v\n", valid, passInfo)
	}
	fmt.Printf("\n\nTotal passports:%d, valid passports:%d, invalid passports:%d\n", numPass, validPass, invalidPass)
	fmt.Printf("\n\nTotal passports:%d, super valid passports:%d,super invalid passports:%d\n", numPass, superValid, superInvalid)
}

type passInfo struct {
	ecl string
	pid string
	eyr string
	hcl string
	byr string
	iyr string
	cid string
	hgt string
}

func parsePassports(lines []string) []*passInfo {
	var retVal = make([]*passInfo, 0)
	var curPass *passInfo
	var re, _ = getPropertyRegex()
	for _, line := range lines {
		var props = getLineProperties(re, line)
		if len(props) > 0 {
			if curPass == nil {
				curPass = new(passInfo)
				retVal = append(retVal, curPass)
			}
			for _, rawP := range props {
				curPass.addProp(&rawP)
			}
		} else {
			curPass = nil //empty line finish current passport
		}
	}

	return retVal
}

func (pi *passInfo) addProp(rp *rawProp) bool {
	switch rp.name {
	case "ecl":
		pi.ecl = rp.value
	case "pid":
		pi.pid = rp.value
	case "eyr":
		pi.eyr = rp.value
	case "hcl":
		pi.hcl = rp.value
	case "byr":
		pi.byr = rp.value
	case "iyr":
		pi.iyr = rp.value
	case "cid":
		pi.cid = rp.value
	case "hgt":
		pi.hgt = rp.value
	default:
		return false
	}
	return true
}

// A passport isValid if it has all fields completed except "cid" which is optional
// returns (validFlag, errorDescription)
// errorDescription a string with the missing fields or "" if valid
func (pi *passInfo) isValid() (bool, string) {
	var errors = make([]string, 0)

	if pi.ecl == "" {
		errors = append(errors, "ecl")
	}
	if pi.pid == "" {
		errors = append(errors, "pid")
	}
	if pi.eyr == "" {
		errors = append(errors, "eyr")
	}
	if pi.hcl == "" {
		errors = append(errors, "hcl")
	}
	if pi.iyr == "" {
		errors = append(errors, "iyr")
	}
	if pi.byr == "" {
		errors = append(errors, "byr")
	}
	if pi.hgt == "" {
		errors = append(errors, "hgt")
	}

	var errString string
	if len(errors) > 0 {
		errString = "[" + strings.Join(errors, ", ") + "]"
	}

	return len(errors) == 0, errString
}

func getYear(str string) (int64, error) {
	var val, err = strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return val, err
}

var heightRe, _ = regexp.Compile(`^(\d+)(cm|in)$`)

func getHeight(str string) (int64, string, error) {
	var props = heightRe.FindStringSubmatch(str)
	if props == nil || len(props) != 3 {
		return 0, "", errors.New("invalid height")
	}

	var heightStr = props[1]
	var height, err = strconv.ParseInt(heightStr, 10, 64)

	if err != nil {
		return 0, "", err
	}

	var unit = props[2]
	return height, unit, nil
}
func isValidHeight(str string) (bool, error) {
	var height, unit, err = getHeight(str)

	if err != nil {
		return false, err
	}

	if unit == "cm" && height >= 150 && height <= 193 {
		return true, nil
	}
	if unit == "in" && height >= 59 && height <= 76 {
		return true, nil
	}

	return false, errors.New("invalid height:" + str)
}

var colorRe, _ = regexp.Compile(`^#[0-9a-f]{6}$`)

func isValidHairColor(str string) (bool, error) {
	var found = colorRe.FindString(str)
	if found == "" {
		return false, errors.New("Invalid color:" + str)
	}
	return true, nil
}

var eyeColorRe, _ = regexp.Compile("^amb|blu|brn|gry|grn|hzl|oth$")

func isValidEyeColor(str string) (bool, error) {
	var found = eyeColorRe.FindString(str)
	if found == "" {
		return false, errors.New("Invalid eye color:" + str)
	}
	return true, nil
}

var idRe, _ = regexp.Compile(`^\d{9}$`)

func isValidPassId(str string) (bool, error) {
	var found = idRe.FindString(str)
	if found == "" {
		return false, errors.New("Invalid password id:" + str)
	}
	return true, nil
}

func (pi *passInfo) isSuperValid() (bool, string) {
	var errs = make([]string, 0)
	var year int64
	var err error

	year, err = getYear(pi.byr)
	if err != nil || year < 1920 || year > 2002 {

		errs = append(errs, "byr")
	}

	year, err = getYear(pi.iyr)
	if err != nil || year < 2010 || year > 2020 {
		errs = append(errs, "iyr")
	}

	year, err = getYear(pi.eyr)
	if err != nil || year < 2020 || year > 2030 {
		errs = append(errs, "eyr")
	}

	_, err = isValidHeight(pi.hgt)
	if err != nil {
		errs = append(errs, "hgt"+err.Error())
	}

	_, err = isValidHairColor(pi.hcl)
	if err != nil {
		errs = append(errs, "hcl")
	}

	_, err = isValidEyeColor(pi.ecl)
	if err != nil {
		errs = append(errs, "ecl")
	}

	_, err = isValidPassId(pi.pid)
	if err != nil {
		errs = append(errs, "pid")
	}

	var errString string
	if len(errs) > 0 {
		errString = "[" + strings.Join(errs, ", ") + "]"
	}

	return len(errs) == 0, errString
}

func getPropertyRegex() (*regexp.Regexp, error) {
	const property = `([a-zA-Z]*):([#a-zA-Z0-9]+)`
	return regexp.Compile(property)
}

type rawProp struct {
	name  string
	value string
}

func getLineProperties(re *regexp.Regexp, line string) []rawProp {
	var props = re.FindAllStringSubmatch(line, -1)
	var retVal = make([]rawProp, 0)
	for _, prop := range props {
		var val = rawProp{name: prop[1], value: prop[2]}
		retVal = append(retVal, val)
	}
	return retVal
}
