package main

import "fmt"

//添加特定的空变量给接口，以保证类型总是实现了预期的接口，避免编译失败
var _ Door = &GlassDoor{}

type GlassDoor struct {
}

func (*GlassDoor) Open() {
	fmt.Println("GlassDoor open")
}

func (*GlassDoor) Close() {
	fmt.Println("GlassDoor close")
}
