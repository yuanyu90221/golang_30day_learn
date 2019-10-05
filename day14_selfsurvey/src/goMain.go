package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {
	names := [...]string{
		"Mon", "Tue", "Wed", "Thu", "Fri",
	}
	weeks := [...]string{
		"Mon", "Tue", "Wed", "Thu", "Fri",
	}
	fmt.Println("names == weeks", names == weeks)

	a := 1
	b := 2
	defer calc("1", a, calc("10", 2, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
	defer func() {
		fmt.Println("ironman")
		fmt.Println("Day 13 post success")
	}()
}
