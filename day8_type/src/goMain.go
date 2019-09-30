package main

import (
	"fmt"

	"day8_type/pointer"
)

type name = string

type newStr string

func SayName(str name) {
	fmt.Println(str)
}

func Say(str newStr) {
	fmt.Println(str)
}
func main() {
	var str = "test"
	SayName(str)

	var ns newStr
	ns = "test newStr"
	Say(ns)
	p := pointer.New(10, 20)
	fmt.Println("p: ", p)

	demoBag()
	ptrDemo()
}

type Bag struct {
	items []int
}

func Insert(b *Bag, itemId int) {
	fmt.Printf("address of *b: %p\n", b)
	b.items = append(b.items, itemId)
}

func InsertValue(b Bag, itemId int) Bag {
	fmt.Printf("address of b: %p\n", &b)
	b.items = append(b.items, itemId)
	return b
}

func demoBag() {
	bag := new(Bag)
	fmt.Printf("address of bag: %p\n", bag)
	fmt.Println("新增元素前給ptr: ", bag)
	Insert(bag, 1000)

	fmt.Println("新增元素後給ptr: ", bag)

	bagValue := Bag{}
	fmt.Printf("address of bagValue: %p\n", bag)
	fmt.Println("新增元素前給實例前: ", bagValue)
	InsertValue(bagValue, 1001)
	fmt.Println("新增元素後, 但沒賦值回去: ", bagValue)
	bagValue = InsertValue(bagValue, 1001)
	fmt.Println("新增元素後, 有沒賦值回去: ", bagValue)
}

func PrintMap(m map[string]int) {
	fmt.Printf("address of map: %p\n", m)
}

func PrintFunc(f func()) {
	fmt.Printf("address of func: %p\n", f)
}

func PrintChan(c chan int) {
	fmt.Printf("address of chan: %p\n", c)
}

func PrintSlice(s []int) {
	fmt.Printf("address of slice: %p\n", s)
}

func PrintArray(a [3]int) {
	fmt.Printf("address of array: %p\n", a)
}

func PrintArrayPtr(a *[3]int) {
	fmt.Printf("address of array: %p\n", a)
}

func ptrDemo() {
	m := make(map[string]int)
	fmt.Printf("address of map: %p\n", m)
	PrintMap(m)

	fun := func() {
		fmt.Println("func")
	}
	fmt.Printf("address of func: %p\n", fun)
	PrintFunc(fun)

	channel := make(chan int)
	fmt.Printf("address of chan: %p\n", channel)
	PrintChan(channel)

	s := make([]int, 10)
	fmt.Printf("address of slice: %p\n", s)
	PrintSlice(s)

	a := [3]int{1, 2, 3}
	fmt.Printf("value of array: %p\n", a)
	fmt.Printf("address of array: %p\n", &a)
	PrintArray(a)
	PrintArrayPtr(&a)
}
