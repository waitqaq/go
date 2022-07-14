package main

import "fmt"

func main() {
	fmt.Println(fibonacci(10))
}

func fibonacci(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
