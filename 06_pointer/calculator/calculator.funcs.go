package main

func (c *Calculator) Add() int {
	tempResult := c.left + c.right
	c.result = tempResult
	return tempResult
}

func (c Calculator) Sub() int {
	return 0
}
func (c Calculator) Multiple() int {
	return 0
}
func (c Calculator) Divide() int {
	return 0
}
func (c Calculator) Reminder() int {
	return 0
}
