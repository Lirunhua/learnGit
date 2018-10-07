package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 55
	ch <- 89
	x := <-ch
	y := <-ch
	fmt.Println(x,y)
}