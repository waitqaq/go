package main

import "fmt"

func main() {
	a, b := 1, 2
	// & 获取变量地址
	// 传入 a b 的指针
	fmt.Println(&a, &b)
	add(&a, &b)
	fmt.Println(a)

	c := &a // c 的类型是 *int，c 指向 a 的盒子，*c 拿到 a 里面的东西
	d := &c // d 的类型是 **int，指向 c 的盒子，本身是指针
	fmt.Println(c)
	fmt.Println("d=", d, "*d=", *d, "**d=", **d)

	m := map[string]string{}
	mp1 := &m
	fmt.Println(mp1)
	put(m)
	fmt.Println(*mp1)

	f1 := add // f1 = func(int, int)
	f1(&a, &b)
	fmt.Println("f1, add = ", a)
	f1p1 := &f1     // f1p1 = *func(int, int)
	(*f1p1)(&a, &b) // *f1p1 先得到函数本身
	fmt.Println("f1p1, add = ", a)

	var array []int
	fmt.Println(array)

}

func put(m map[string]string) {
	m["a"] = "AAA"
}

// *int 表示变量为 int 的指针
func add(a, b *int) {
	fmt.Println(*a, *b)
	*a = *a + *b
}
