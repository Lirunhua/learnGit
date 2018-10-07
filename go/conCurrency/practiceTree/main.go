package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch, 中根遍历
// tree.Tree is a type 


/* MY WRONG SOLUTINO
func Walk(t *tree.Tree, ch chan int){
	if t.Left != nil{
		fmt.Println("left is not nil")
		Walk(t.Left, ch)
	}else{
		fmt.Println("reading value")
		ch <- t.Value
		if t.Right != nil {
			Walk(t.Right, ch)
		}
	}
		
}
*/

// right solution
func Walk(t *tree.Tree, ch chan int){
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}



// Same 检测树 t1 和 t2 是否含有相同的值。
//func Same(t1, t2 *tree.Tree) bool

func main() {
	ch1 := make(chan int,10)
	ch2 := make(chan int,10)
	a := tree.New(3) // type(a) = tree.Tree
	b := tree.New(2)
	go Walk(a, ch1)
	go Walk(b, ch2)
	for i:=0; i<10; i++{
		fmt.Println(<-ch2)
	}


	for i:=0; i<10; i++{
		fmt.Println(<-ch1)
	}
}
