package extend

import "fmt"

func init() {
	fmt.Println("extend ==> init()")
}

func Min(a float64, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}
