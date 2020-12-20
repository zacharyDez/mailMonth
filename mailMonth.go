package main

import (
	"errors"
	"strconv"
)

func getInvalidMonthErr(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid month: " + strconv.Itoa(month))
	}

	return nil
}

func isMailMonthInInterval(mailMonth int, startMonth int, endMonth int) (bool, error) {
	err := getInvalidMonthErr(mailMonth)
	if err != nil {
		return false, err
	}

	err = getInvalidMonthErr(startMonth)
	if err != nil {
		return false, err
	}

	err = getInvalidMonthErr(endMonth)
	if err != nil {
		return false, err
	}

	if startMonth > endMonth && (mailMonth >= startMonth || mailMonth <= endMonth) {
		return true, err
	}

	if startMonth <= endMonth && mailMonth >= startMonth && mailMonth <= endMonth {
		return true, err
	}

	return false, err
}
