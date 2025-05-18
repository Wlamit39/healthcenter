package actions

import (
	"time"
)

func CalculateAgeFromDOB(dob string) (int, error) {
	// Parse the date
	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return 0, err
	}

	// Calculate the age
	now := time.Now()
	age := now.Year() - parsedDOB.Year()
	if now.YearDay() < parsedDOB.YearDay() {
		age-- // Birthday hasn’t occurred yet this year
	}

	return age, nil
}
