package main

import "fmt"

type inputFromStd struct {
}

func (inputFromStd) GetInput() Person {
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

	return Person{
		weight: weight,
		tall:   tall,
		age:    age,
		sex:    sex,
	}
}
