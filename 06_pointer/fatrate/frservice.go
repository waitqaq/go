package main

import "log"

type fatRateService struct {
	input InputService
	s     *fatRateSuggestion
}

func (frsvc *fatRateService) GiveSuggestionToPerson(person *Person) string {
	if err := person.calcBmi(); err != nil {
		log.Printf("无法给出 %s 计算 BMI:%v", person.name, err)
		return "error"
	}
	person.calcFatRate()
	return frsvc.s.GetSuggestion(person)
}

func (frsvc *fatRateService) GiveSuggestionToPersons(persons ...*Person) map[*Person]string {
	out := map[*Person]string{}
	for _, item := range persons {
		out[item] = frsvc.GiveSuggestionToPerson(item)
	}
	return out
}
