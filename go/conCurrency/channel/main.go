package main

import "fmt"
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1,2,3,4,5,6,7,8,9}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	/*
	this can be seperated by two processing
	first : x <- c, the c here is given by goroutine "go sum(s[:len(s)/2], c)"
	after the content in c was taken by "x <- c", which given by goroutine 
	"go sum(s[:len(s)/2], c)", then the goroutine "go sum(s[len(s)/2:], c)"
	will send the result to channel c; and 
	"second: y <- c " will take the content from channel c;
	*/

	fmt.Println(x, y, x+y)
}