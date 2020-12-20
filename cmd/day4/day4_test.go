package main

import (
	"testing"
)

func TestGetYearOk(t *testing.T) {
	var goodYears = []string{"123", "1234"}

	for _, y := range goodYears {
		var _, err = getYear(y)
		if err != nil {
			t.Fail()
		}
	}
}

func TestGetYearFail(t *testing.T) {
	var badYears = []string{"12a3", "1234a"}

	for _, y := range badYears {
		var _, err = getYear(y)
		if err == nil {
			t.Fail()
		}
	}

}

func TestIsValidHeight(t *testing.T) {
	var validHeights = []string{"150cm", "160cm", "190cm", "193cm", "59in", "60in", "76in"}

	for _, h := range validHeights {
		var valid, _ = isValidHeight(h)
		if !valid {
			t.Error(h, " is an invalid height")
		}
	}
	var invalidHeights = []string{"149cm", "10cm", "1190cm", "194cm", "58in", "88in", "77in", "3km", "hello"}

	for _, h := range invalidHeights {
		var valid, _ = isValidHeight(h)
		if valid {
			t.Error(h, " is a valid height although is shouldn't be")
		}
	}

}
