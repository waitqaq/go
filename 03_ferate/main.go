package main

import (
	"fmt"
	"learning/03_ferate/calc"
)

func main() {
	for {
		mainFateRateBody()
		if cont := whetherContinue(); !cont {
			break
		}
	}
}

func mainFateRateBody() {
	weight, tall, age, _, sex := getMaterialsFromInput()
	fateRate := calcFateRate(weight, tall, age, sex)
	fmt.Println(fateRate)
}

func calcFateRate(weight float64, tall float64, age int, sex string) float64 {
	bmi := calc.CalcBMI(tall, weight)
	fateRate := calc.CalcFatRate(bmi, age, sex)
	return fateRate
}

func getMaterialsFromInput() (float64, float64, int, int, string) {
	var weight float64
	fmt.Print("体重（kg）：")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高（m）：")
	fmt.Scanln(&tall)

	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	var sexWeight int
	sex := "男"
	fmt.Print("性别（男/女）:")
	fmt.Scanln(&sex)

	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	return weight, tall, age, sexWeight, sex
}

func whetherContinue() bool {
	var whetherContinue string
	fmt.Print("是否录入下一个（y/n）？")
	fmt.Scanln(&whetherContinue)
	if whetherContinue != "y" {
		return false
	}
	return true
}
