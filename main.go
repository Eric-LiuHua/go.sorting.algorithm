package main

import (
	"fmt"
	"go_sorts/sorts"
)

func main() {
	fmt.Println("go-sorts test start !!!")
	b1 := []int{7, 89, 99, 55, 2235, 515, 88, 22, 2}
	fmt.Println("b1 org:", b1)
	//b1 = sorts.BubbleSort(b1[:])
	b1 = sorts.ShellSortA(b1[:])
	fmt.Println("b1 org:", b1)

}
