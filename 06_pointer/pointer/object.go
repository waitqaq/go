package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)

func main() {

	person := []Person{
		{
			"小强",
			"男",
			1.7,
			70,
			35,
			18,
		},
	}
	for _, item := range person {
		bmi, err := gobmi.BMI(item.weight, item.height)
		fmt.Println(bmi, err)
	}
	a := new(Person)
	b := *a
	fmt.Println(b)
}

type Person struct {
	name   string
	sex    string
	tall   float64
	weight float64
	height float64
	age    int
}
