package main

import "fmt"

func main() {
	//person := getFakePersonInfo()
	//c := Calc{}
	//c.BMI(person)
	//c.FatRate(person)
	//fmt.Println(person)
	//sug := fatRateSuggestion{}
	//suggestion := sug.GetSuggestion(person)
	//fmt.Println(suggestion)

	frSvc := &fatRateService{s: GetFatRateSuggestion(), input: &inputFromStd{}}

	for {
		p := frSvc.input.GetInput()
		fmt.Println(frSvc.GiveSuggestionToPerson(&p))
	}

}
