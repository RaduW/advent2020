package main

import (
	"advent2020/internal/util"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type passwordInfo struct {
	min    uint64
	max    uint64
	sign   rune
	passwd string
}

func main() {
	fmt.Println("hello")
	var passwords, err = util.GetLines("passwords.txt")
	if err != nil {
		fmt.Println("Error importing passwords")
		return
	}
	var passwdPolicy *regexp.Regexp
	passwdPolicy, err = getPasswordPolicy()
	if err != nil {
		fmt.Println("Error compiling password policy")
		return
	}

	var validPasswords = 0
	var invalidPasswords = 1
	var totalPasswords = len(passwords)
	for _, passwd := range passwords {
		var matchInfo, _ = getPasswordInfo(passwdPolicy.FindStringSubmatch(passwd))
		var matches = validatePasswordInfo(matchInfo)
		var matchStr string
		if matches {
			matchStr = "matches"
			validPasswords += 1
		} else {
			matchStr = "does NOT match"
			invalidPasswords += 1
		}
		fmt.Println("passwd: ", passwd, " ", matchStr)

	}

	fmt.Printf("\nFro %d passwords %d match and %d don't \n", totalPasswords, validPasswords, invalidPasswords)

}

func validatePasswordInfo(info *passwordInfo) bool {
	var cnt uint64 = 0
	for _, char := range info.passwd {
		if char == info.sign {
			cnt += 1
		}
	}
	return cnt >= info.min && cnt <= info.max
}

func getPasswordInfo(parsedLine []string) (*passwordInfo, error) {
	var err error
	var min, max uint64

	if len(parsedLine) != 5 {
		return nil, errors.New("invalid parsed line length")
	}
	var l = parsedLine
	var minStr, maxStr, signStr, passwd = l[1], l[2], l[3], l[4]
	min, err = strconv.ParseUint(minStr, 10, 32)
	if err != nil {
		return nil, err
	}
	max, err = strconv.ParseUint(maxStr, 10, 32)
	if err != nil {
		return nil, err
	}

	var sign rune
	for _, r := range signStr {
		sign = r
		break
	}

	return &passwordInfo{min, max, sign, passwd}, nil
}

func getPasswordPolicy() (*regexp.Regexp, error) {
	const passwordPolicy = `(\d+)-(\d+)\s*(\S):\s*(\S+)`
	return regexp.Compile(passwordPolicy)
}
