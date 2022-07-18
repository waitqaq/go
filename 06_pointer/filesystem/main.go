package main

import (
	"fmt"
)

func main() {
	var equipment IOInterface = &Soft{}
	data := equipment.Read()
	fmt.Println(data)
}

type IOInterface interface {
	Read() string
}

type Soft struct {
}

func (Soft) Read() string {
	return "软"
}
