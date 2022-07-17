package main

import (
	"fmt"
)

func main() {
	var left, right int = 1, 2
	c := Calculator{
		left:  left,
		right: right,
	}
	fmt.Println(c.Add())
	fmt.Println(c.result)

	newC := NewCalculator{}
	newC.left = 100
	newC.right = 200
	fmt.Println(newC.Add())

	newP := pointerCalculator{&Calculator{}}
	newP.left = 200
	newP.right = 300
	fmt.Println(newP.Add())

	mc := MyCommand{
		map[string]string{},
	}
	mc.commandOptions["aa"] = "AAA"
	fmt.Println(mc.ToCmdStr())
}

type Calculator struct {
	left, right int
	op          string
	result      int
}

type MyCommand struct {
	//mainCommand    *string
	commandOptions map[string]string
}

func (my MyCommand) ToCmdStr() string {
	out := ""
	for k, v := range my.commandOptions {
		out = out + fmt.Sprintf("--%s=%s", k, v)
	}
	return out
}
