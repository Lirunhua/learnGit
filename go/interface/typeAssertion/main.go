package main

import(
	"fmt"
)

func main() {
	var i interface{} = "Hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)
	
	f, ok := i.(float64)  // this will be ok and won't triger panic
	fmt.Println(f, ok)
	
	te := i.(float64) // this will triger panic
	fmt.Println(te)
}