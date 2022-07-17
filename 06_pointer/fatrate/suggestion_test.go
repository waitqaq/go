package main

import "testing"

func Test_fatRateSuggestion_GetSuggestion(t *testing.T) {
	sug := GetFatRateSuggestion()
	tests := []Person{
		{
			age:     35,
			sex:     "男",
			fatRate: 0.24,
		},
	}
	if got := sug.GetSuggestion(&tests[0]); got != "非常肥胖" {
		t.Fail()
	}
}
