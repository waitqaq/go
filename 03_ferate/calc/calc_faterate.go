package calc

import (
	gobmi "github.com/armstrongli/go-bmi"
)

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64) {
	return gobmi.CalcFatRate(bmi, age, sex)
}

func getAgeWeight(age int) (ageWeight float64) {
	ageWeight = 0.23
	if age >= 30 && age <= 40 {
		ageWeight = 0.22
	}
	return
}
