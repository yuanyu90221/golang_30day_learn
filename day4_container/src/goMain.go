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
	t = time.Now()
	l := list.New()
	for i := 0; i < 1*100000*1000; i++ {
		l.PushBack(1)
	}
	fmt.Println("List 新增元素耗費: " + time.Now().Sub(t).String())

	t = time.Now()
	for e := l.Front(); e != nil; e = e.Next() {
	}
	fmt.Println("走訪List耗費:" + time.Now().Sub(t).String())

	var em *list.Element
	i := 0
	// 找到1/3處的元素
	for e := l.Front(); e != nil; e = e.Next() {
		i++
		if i == l.Len()/3 {
			em = e
			break
		}
	}
	// 因為是記算插入元素的速度, 所以忽略查找的時間
	t = time.Now()
	l.InsertAfter(2, em)
	fmt.Println("List 的插入元素耗費 : " + time.Now().Sub(t).String())
	// 比較刪除元素
	t = time.Now()
	l.Remove(em)
	fmt.Println("List 的刪除元素耗費:" + time.Now().Sub(t).String())
}
