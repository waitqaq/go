package main

import "fmt"

type WoodDoor struct {
}

func (*WoodDoor) Open() {
	fmt.Println("GlassWood Open")
}

func (*WoodDoor) Close() {
	fmt.Println("GlassWood Close")
}
