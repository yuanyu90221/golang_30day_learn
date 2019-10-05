package main

import (
	. "day10_package/math"
	"fmt"
)

func init() {
	fmt.Println("main ==> init()")
}
func main() {
	fmt.Println("hello")
	fmt.Println(Average([]float64{1, 2}))
	// fmt.Println(Min(1, 2))
}
