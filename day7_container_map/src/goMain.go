package main

import (
	"fmt"
	"sync"
)

func main() {
	mapDemo := map[int]int{
		1:  1,
		2:  2,
		3:  3,
		4:  4,
		5:  5,
		6:  6,
		7:  7,
		8:  8,
		9:  9,
		10: 10,
	}
	for k, v := range mapDemo {
		fmt.Println("value: ", v, "key: ", k)
	}
	dictDemo()
	syncDemo()
}

func dictDemo() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println("map: ", n)
}

var printMap = func(key, value interface{}) bool {
	fmt.Printf("key: %s, value: %d\n", key, value)
	return true
}

func syncDemo() {
	var m sync.Map

	m.Store("k1", 7)
	m.Store("k2", map[string]int{"k4": 5})

	m.Range(printMap)
	fmt.Println("----------")

	v1, _ := m.Load("k1")
	fmt.Println("v1: ", v1)

	fmt.Println("----------")
	m.Delete("k2")

	v1, exist := m.LoadOrStore("k1", 8)
	fmt.Printf("v1: %d, exist: %v\n", v1, exist)
	v3, exist3 := m.LoadOrStore("k3", 2)
	fmt.Printf("v3: %d, exist: %v\n", v3, exist3)
	m.Range(printMap)
	fmt.Println("------------")

	_, exist2 := m.Load("k2")
	fmt.Printf("v3  exist: %v\n", exist2)
}
