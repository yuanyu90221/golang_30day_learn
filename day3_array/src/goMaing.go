package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	arrA := [2]int{}
	var arrB [2]int
	arrC := [2]int{1, 2}
	arrB = arrA

	fmt.Printf("arrA : %p, %v\n", &arrA, arrA)
	fmt.Printf("arrB : %p, %v\n", &arrB, arrB)

	arr(arrA)
	fmt.Printf("arrC: %p, %v\n", &arrC, arrC)
	arrRef(&arrC)
	arrC[0]++
	fmt.Printf("arrC: %p, %v\n", &arrC, arrC)
	reflTest()
	numListDemo()
	arrmake()
	appendDemo()
	appendDemo2()
	appendDemo3()
	appendDemo4()
}

func arr(x [2]int) {
	fmt.Printf("pass Array: %p, %v\n", &x, x)
}

func arrRef(x *[2]int) {
	fmt.Printf("pass Array : %p , %v\n", x, *x)
	time.Sleep(time.Second)
	(*x)[0]++
}

func reflTest() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice0 := arr[:5]

	fmt.Println(slice0)

	slice1 := arr[5:]
	fmt.Println(slice1)

	slice2 := arr[2:7]
	fmt.Println(slice2)

	slice3 := arr[:]
	fmt.Println(slice3)

	slice4 := arr[0:0]
	fmt.Println(slice4)
	fmt.Println("---------------")
	slice := arr[1:3]
	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(len(slice), cap(slice))

	fmt.Println(arr)
	fmt.Println(slice)

	fmt.Println("-------------------")
	slice[0] = 0
	fmt.Println(arr)
	fmt.Println(slice)

}

func numListDemo() {
	var numList []int

	var numEmptyList = []int{}

	fmt.Println(numList, numEmptyList)

	fmt.Println(len(numList), len(numEmptyList))

	fmt.Println(numList == nil)
	fmt.Println(numEmptyList == nil)
}

func arrmake() {
	a := make([]int, 2)
	b := make([]int, 2, 10)

	fmt.Println(a, b)
	fmt.Println(len(a), len(b))
	fmt.Println(cap(a), cap(b))
}

func appendDemo() {
	numbers := make([]int, 0)

	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d, cap: %d, ptr: %p\n", len(numbers), cap(numbers), numbers)
	}
}

func appendDemo2() {
	a := make([]int, 0, 10)
	b := append(a, 1)
	_ = append(a, 2)
	fmt.Println(b[0])
}

func appendDemo3() {
	a := make([]int, 10, 20)
	b := a[5:]
	fmt.Println(len(b), cap(b))
}

func doAppend(a []int) {
	b := append(a, 0)
	fmt.Println(b)
}

func appendDemo4() {
	a := []int{1, 2, 3, 4, 5}
	doAppend(a[0:2:2])
	fmt.Println(a)
}
