package main

import "fmt"

func main() {
	var refrigerator Refrigerator
	fmt.Println(refrigerator)
}

type Open interface {
	Open() error
}

type Close interface {
	Close() error
}

type Box interface {
	Open
	Close
}

type Refrigerator struct {
	Size string
}

func (Refrigerator) Open() error {
	return nil
}

func (Refrigerator) Close() error {
	return nil
}
