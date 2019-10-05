package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	defer fmt.Println("defer main")
	var user = os.Getenv("USER_")
	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("error:", err)
				fmt.Println("recover success.")
			}
		}()

		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic("should set user env.")
			}
			fmt.Println("after panic")
		}()
	}()

	time.Sleep(1 * time.Second)
	result := 1
	fmt.Printf("get result %d\r\n", result)
	deferLoop()
}

func deferLoop() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
