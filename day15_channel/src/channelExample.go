package main

import (
	"fmt"
)

func printer(c chan int) {
	for {
		data := <-c
		if data == 0 {
			fmt.Println("break")
			break
		}
		fmt.Println(data)
	}
	// notify finish
	c <- 0
}
func main() {
	c := make(chan int)

	go printer(c)

	for i := 1; i <= 10; i++ {
		c <- i
	}

	c <- 0

	<-c
}
