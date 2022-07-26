package main

import (
	"fmt"
	"testing"
)

type Change interface {
	ChangeName(newName string)
	ChangeAge(newAge int)
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) ChangeName(newName string) {
	s.Name = newName
}

func (s *Student) ChangeAge(newAge int) {
	s.Age = newAge
}

func TestVal(t *testing.T) {
	var stdChg Change
	// stdChg = Student   // 当实现接口得方法是在对象指针上，只能用对象指针作为值赋值给接口
	stdChg = &Student{
		Name: "小明",
		Age:  18,
	}
	fmt.Println(stdChg)

}

func TestStd(t *testing.T) {
	s := Student{Name: "小明"}
	s.ChangeName("小李")
	fmt.Println(s.Name)
}
