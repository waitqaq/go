package main

import (
	"fmt"
	"learning/03_ferate/calc"
	"runtime/debug"
)

func main() {
	for {
		mainFateRateBody()
		if cont := whetherContinue(); !cont {
			break
		}
	}
}

func recoverMainBody() {
	if re := recover(); re != nil {
		fmt.Printf("warning: catch critical error: %v \n", re)
	}
	debug.PrintStack()
}

func mainFateRateBody() {
	defer recoverMainBody() // 捕获异常错误
	weight, tall, age, _, sex := getMaterialsFromInput()
	fateRate, err := calcFateRate(weight, tall, age, sex)
	if err != nil {
		fmt.Println("warning :计算体脂率出错", err)
		return
	}
	fmt.Println(fateRate)
}

func calcFateRate(weight float64, tall float64, age int, sex string) (fateRate float64, err error) {
	// 计算体脂率
	bmi, err := calc.CalcBMI(tall, weight)
	if err != nil {
		return 0, err
	}
	fateRate = calc.CalcFatRate(bmi, age, sex)
	return fateRate, err
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
