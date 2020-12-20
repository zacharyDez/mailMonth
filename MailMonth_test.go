package main

import (
	"testing"
)

func TestMailMonthInvalid(t *testing.T) {
	startMonth := 12
	endMonth := 5

	mailMonth := 13
	_, err := isMailMonthInInterval(mailMonth, startMonth, endMonth)

	if err == nil {
		t.Errorf("Expected Start Month (%d) to be invalid.", mailMonth)
	}

	startMonth = 12
	endMonth = 5

	mailMonth = 0
	_, err = isMailMonthInInterval(mailMonth, startMonth, endMonth)

	if err == nil {
		t.Errorf("Expected Start Month (%d) to be invalid.", mailMonth)
	}
}

func TestMailMonthStartInvalid(t *testing.T) {
	startMonthGreater := 13
	endMonth := 5

	monthInInterval := 2
	_, err := isMailMonthInInterval(monthInInterval, startMonthGreater, endMonth)

	if err == nil {
		t.Errorf("Expected Start Month (%d) to be invalid.", startMonthGreater)
	}

	startMonthLower := 0
	endMonth = 5

	monthInInterval = 2
	_, err = isMailMonthInInterval(monthInInterval, startMonthLower, endMonth)

	if err == nil {
		t.Errorf("Expected Start Month (%d) to be invalid.", startMonthLower)
	}
}

func TestMailMonthEndInvalid(t *testing.T) {
	startMonth := 12
	endMonthGreater := 13

	monthInInterval := 2
	_, err := isMailMonthInInterval(monthInInterval, startMonth, endMonthGreater)

	if err == nil {
		t.Errorf("Expected End Month (%d) to be invalid.", endMonthGreater)
	}

	startMonth = 1
	endMonthLower := 0

	monthInInterval = 2
	_, err = isMailMonthInInterval(monthInInterval, startMonth, endMonthLower)

	if err == nil {
		t.Errorf("Expected Start Month (%d) to be invalid.", endMonthLower)
	}
}

func TestMailMonthStartGreaterWithin(t *testing.T) {
	startMonth := 10
	endMonth := 5

	monthInInterval := 2
	isIncluded, _ := isMailMonthInInterval(monthInInterval, startMonth, endMonth)

	if isIncluded != true {
		t.Errorf("Expected month to be included in interval, got: %v, want: %v.", isIncluded, true)
	}
}


func TestMailMonthStartSmallerWithin(t *testing.T) {
	startMonth := 5
	endMonth := 10

	monthInInterval := 7
	isIncluded, _ := isMailMonthInInterval(monthInInterval, startMonth, endMonth)

	if isIncluded != true {
		t.Errorf("Expected month to be included in interval, got: %v, want: %v.", isIncluded, true)
	}
}

func TestMailMonthStartSmallerEqualStart(t *testing.T) {
	startMonth := 5
	endMonth := 10

	monthInInterval := 5
	isIncluded, _ := isMailMonthInInterval(monthInInterval, startMonth, endMonth)

	if isIncluded != true {
		t.Errorf("Expected month to be included in interval, got: %v, want: %v.", isIncluded, true)
	}
}

func TestMailMonthStartSmallerEqualEnd(t *testing.T) {
	startMonth := 5
	endMonth := 10

	monthInInterval := 10
	isIncluded, _ := isMailMonthInInterval(monthInInterval, startMonth, endMonth)

	if isIncluded != true {
		t.Errorf("Expected month to be included in interval, got: %v, want: %v.", isIncluded, true)
	}
}


func TestMailMonthStartSmallerOut(t *testing.T) {
	startMonth := 5
	endMonth := 10

	monthInInterval := 11
	isIncluded, _ := isMailMonthInInterval(monthInInterval, startMonth, endMonth)

	if isIncluded != false {
		t.Errorf("Expected month to be included in interval, got: %v, want: %v.", isIncluded, false)
	}
}

func TestMailMonthStartGreaterOut(t *testing.T) {
	startMonth := 10
	endMonth := 5

	monthInInterval := 7
	isIncluded, _ := isMailMonthInInterval(monthInInterval, startMonth, endMonth)

	if isIncluded != false {
		t.Errorf("Expected month to be included in interval, got: %v, want: %v.", isIncluded, false)
	}
}


func TestMailMonthStartEqualIn(t *testing.T) {
	startMonth := 10
	endMonth := 10

	monthInInterval := 10
	isIncluded, _ := isMailMonthInInterval(monthInInterval, startMonth, endMonth)

	if isIncluded != true {
		t.Errorf("Expected month to be included in interval, got: %v, want: %v.", isIncluded, true)
	}
}


func TestMailMonthStartEqualOut(t *testing.T) {
	startMonth := 10
	endMonth := 10

	monthInInterval := 1
	isIncluded, _ := isMailMonthInInterval(monthInInterval, startMonth, endMonth)

	if isIncluded != false {
		t.Errorf("Expected month to be included in interval, got: %v, want: %v.", isIncluded, false)
	}
}