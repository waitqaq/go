package main

import "fmt"

func main() {
	for {
		var weight float64
		fmt.Print("体重（kg）：")
		fmt.Scanln(&weight)
		var tall float64 = 1.70
		var bmi float64 = weight / (weight * tall)
		var age int = 35
		var sexWeight int
		var sex string
		if sex == "男" {
			sexWeight = 1
		} else {
			sexWeight = 0
		}
		var fatRate float64 = 1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)
		fmt.Println("体脂率是:", fatRate)
		var whetherContinue string
		fmt.Print("是否录入下一个（y/n）?")
		fmt.Scanln(&whetherContinue)
		if whetherContinue != "y" {
			break
		}
	}

}
