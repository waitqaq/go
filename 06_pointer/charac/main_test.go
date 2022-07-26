package main

import (
	"fmt"
	"testing"
)

func TestAssettion(t *testing.T) {
	r := TestBox{}
	var c Close = r

	switch cDetail := c.(type) {
	case Refrigerator:
		fmt.Println("预期的冰箱")
		fmt.Println(cDetail.Size)
	case TestBox:
		fmt.Println("这是一个box")
	}

}

type TestBox struct {
}

func (tb TestBox) Close() error {
	return nil
}
