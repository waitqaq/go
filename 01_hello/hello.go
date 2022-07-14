package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println("hello world")
	//var age int = 15
	//fmt.Println(age)

	var ho, ver float32 = 3, 4.1
	var sc = ho * ver
	fmt.Println(sc)

	var name string
	fmt.Println("name=", name)

	var int6 uint = math.MaxUint64
	fmt.Println(int6)
	var int7 int = int(int6)
	fmt.Println(int7)

	age := 30
	fmt.Printf("%p\n", &age)
	age = 32
	fmt.Printf("%p\n", &age)

	const pi = 3.1415
	fmt.Println(pi)

	a1 := "hello"
	fmt.Println(reflect.TypeOf(a1))
	a2 := &a1
	fmt.Println(reflect.TypeOf(a2))

	a, b := 100, 31
	fmt.Println(a ^ b)

	fmt.Printf("字符串:%s\n", "小明")

	var money int = 20
	var busy bool = false
	switch money {
	case 20:
		fmt.Println("hello")
		// 直接进去下层条件
		fallthrough
	case 200:
		fmt.Println("world")
		if busy {
			break
		} else {
			fmt.Println("hao")
		}
	}
	var int16s int16 = 32726
	var tail interface{} = int16s

	switch tail.(type) {
	case int:
		fmt.Println("int")
	case int16:
		fmt.Println("int16")
	default:
		fmt.Println("not know")
	}
	tail = 1
	fmt.Println(reflect.TypeOf(tail))

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	var ages [5]string = [5]string{"1", "3", "4", "5", "6"}
	fmt.Println(ages)
	var c = [...]int{1, 2, 3, 4}
	c[0] = 11
	fmt.Println(c)
	fmt.Println(len(c))

	// 数组
	newPersonInfos := [...][3]string{
		[3]string{"1", "2", "3"},
		[3]string{"1", "2", "3"},
		[3]string{"1", "2", "3"},
		[3]string{"1", "2", "3"},
	}
	for i, val := range newPersonInfos {
		for _, val1 := range newPersonInfos[i] {
			fmt.Println(val, val1)
		}
	}

	// 切片
	sqInfo := []int{1}
	fmt.Println(sqInfo)
	sqInfo = append(sqInfo, 5, 6, 7)
	fmt.Println(sqInfo)
	// 删除元素
	sqInfo = append(sqInfo[:1], sqInfo[2:]...)
	fmt.Println(sqInfo)

	var str string = "hello"
	strByte := []byte(str)
	fmt.Println(strByte)
	strByte[0] = 'A'
	fmt.Println(string(strByte))

	// 针对中文转 ascall 采用 rune
	var strZh string = "您好"
	strZhByte := []rune(strZh)
	fmt.Println(strZhByte)
	strZhByte[0] = 'A'
	fmt.Println(string(strZhByte))

	makeA := make([]int, 1, 3)
	fmt.Println("len:", len(makeA), "val", makeA, "cap:", cap(makeA))
	makeA = append(makeA, 3)
	fmt.Println("len:", len(makeA), "val", makeA, "cap:", cap(makeA))

	// map
	var mapA map[string]int
	fmt.Println(mapA)
	mapB := map[string]int{"a": 1}
	mapB["b"] = 2
	fmt.Println(mapB)
	delete(mapB, "a")
	fmt.Println(mapB)
	fmt.Println(mapB["b"])
}
