package main

import (
	"fmt"
)

func twice (i interface {}) {
	switch v := i.(type) {
	case int :
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("Twice %v is %v\n", v, v+v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	twice(21)
	twice("hello")
	twice(true)
}