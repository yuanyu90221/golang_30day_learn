package math

import (
	"fmt"
)

func init() {
	fmt.Println("math ==> init()")
}

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// func Min(a float64, b float64) float64 {
// 	return extend.Min(a, b)
// }
