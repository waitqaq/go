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

	frSvc := &fatRateService{s: GetFatRateSuggestion()}

	for {
		p := getPersonInfoFromInput()
		fmt.Println(frSvc.GiveSuggestionToPerson(p))
	}

}

func getPersonInfoFromInput() *Person {
	var weight float64
	fmt.Print("体重（kg）：")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高（m）：")
	fmt.Scanln(&tall)

	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	var sex string
	fmt.Print("性别（男/女）:")
	fmt.Scanln(&sex)

	return &Person{
		weight: weight,
		tall:   tall,
		age:    age,
		sex:    sex,
	}
}

func getFakePersonInfo() *Person {
	return &Person{
		name:   "小强",
		sex:    "男",
		tall:   1.7,
		weight: 70,
		age:    30,
	}
}
