package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//openFile()
	//deferGuess()
	panicSam(-1)
}

func openFile() {
	fileName := "test.txt"
	fileObj, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	defer fileObj.Close()
}

func deferGuess() {
	startTime := time.Now()
	fmt.Println("start time:", startTime)
	// defer 本身只关注对应的函数本身
	//defer fmt.Println("duration:", time.Now().Sub(startTime))
	defer func() {
		fmt.Println("duration:", time.Now().Sub(startTime))
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("end time:", time.Now())
}

func panicSam(age int) {
	defer coverPanic()
	if age < 0 {
		panic("age 不允许为负数")
	}
}

func coverPanic() {
	if r := recover(); r != nil {
		fmt.Println("系统出现错误")
	}
}
