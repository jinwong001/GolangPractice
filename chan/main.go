package main

import (
	"fmt"
)

func main() {
	//c := make(chan int)
	//a := []int{1, 8, 9, 4, 5}
	////sum(a, c)
	//go sum(a[:len(a)/2], c)
	//go sum(a[len(a)/2:], c)
	//x, y := <-c, <-c // 从 c 中获取
	//fmt.Println(x, y, x+y)

	sendChan()
	//fmt.Scanln()
	fmt.Println("done!")
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
	close(c)
}

func sendChan() {
	c := make(chan int)
	//go func() {
	//	c <- 1
	//}()
	c<-1
	v := <-c
	fmt.Println(v)
}
