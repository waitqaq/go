package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	convertType()
}

func convertType() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic啦")
			debug.PrintStack()
		}
	}()
	var a interface{}
	a = "string"
	b := a.(int)
	fmt.Println(b)
}
