package main

import (
	"fmt"
	//"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i<n ; i++{
		c <- x	
		/* if channel is full, this go routine will wait here; until the content
		in the channel was taken by the other process
		*/
		x, y = y, x+y 
	}
	close(c)
}

func main() {
	ch := make(chan int, 3)
	go fibonacci(30, ch)

	
	for i := range ch {
		fmt.Println(i)
	}

	v, ok := <-ch
	fmt.Println(v, ok) // here the v will be 0, not nil; because ch have been initialized
}