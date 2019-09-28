package main

import (
	"container/list"
	"fmt"
	"time"
)

func traverse(list *list.List) {
	fmt.Printf("root -> ")
	for el := list.Front(); el != nil; el = el.Next() {
		fmt.Printf("%v -> ", el.Value)
	}
}

func main() {
	list := list.New()

	list.PushBack(20)

	list.PushFront("10")

	element := list.PushBack(25)

	list.InsertAfter("26", element)

	list.InsertBefore(24, element)

	traverse(list)

	fmt.Println("\n------------")

	list.MoveAfter(element, list.Front())

	traverse(list)

	fmt.Println("\n------------")

	list.MoveBefore(element, list.Front())

	traverse(list)

	fmt.Println("\n------------")

	list.MoveToBack(element)

	traverse(list)

	fmt.Println("\n------------")

	list.MoveToFront(element)

	traverse(list)

	fmt.Println("\n------------")

	list.Remove(element)

	traverse(list)

	fmt.Println("\n------------")

	timeTest()
}

func timeTest() {
	t := time.Now()

	sli := make([]int, 10)

	for i := 0; i < 1*100000*1000; i++ {
		sli = append(sli, 1)
	}
	fmt.Println("Slice cost: ", time.Now().Sub(t).String())
	t = time.Now()
	for _ = range sli {
	}
	fmt.Println("transversal cost: ", time.Now().Sub(t).String())

	t = time.Now()
	slif := sli[:100000*500]
	slib := sli[100000*500:]

	slif = append(slif, 10)
	slif = append(slif, slib...)
	fmt.Println("Slice 的插入元素耗費 : " + time.Now().Sub(t).String())

	// 比較刪除元素
	t = time.Now()
	index := 100000
	_ = append(sli[:index], sli[index+1:]...)
	fmt.Println("Slice 的刪除元素耗費 : " + time.Now().Sub(t).String())

	sli = make([]int, 10)

	// ---------Slice end, start list
	fmt.Println("------------------------------")
}
