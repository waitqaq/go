package main

import (
	"fmt"
	"reflect"
	"time"
)

var NAME string = "a"

func init() {
	fmt.Println("hello")
}
func init() {
	fmt.Println("world")
}
func main() {
	//helloName("li")
	//internal()
	//constructHello()
	sampleName()
}

func Hello() {
	fmt.Println("hello")
}

func sampleName() {
	name := "小强"
	fmt.Println("名字是", name)
	{
		name := "小强 2"
		fmt.Println("更新后的名字是：", name)
	}
}

func hello() {
	fmt.Printf("hello")
}

func helloName(name string) string {
	fmt.Println("你好", name)
	return "你好" + name
}

func internal() {
	startTime := time.Now()
	defer func() {
		finishTime := time.Now()
		fmt.Println(finishTime.Sub(startTime))
	}()
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	arr := make([]string, 3, 4)
	arr[0] = "1"
	fmt.Println("len:", len(arr), "cap:", cap(arr), "default:", arr)

	i := new(int)
	fmt.Println(i)
	fmt.Println(reflect.TypeOf(i))

	arr2 := make([]string, 3, 4)
	copy(arr2, arr)
	fmt.Println(arr2)
}

func constructHello() {
	//bmi := []float64{1.2, 3.222, 4, 3}
	avgBmi := calculateAvg(2, 3)
	fmt.Println(avgBmi)
}

func calculateAvg(bmi ...float64) (s float64) {
	var total float64 = 0
	for _, item := range bmi {
		total += item
	}
	s = total
	return s
}
