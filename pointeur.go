package main

import (
	"fmt"
)

func main() {
	//var a *int
	//zero(&a)
	//fmt.Println(*a)

	var i int = 18 //*int = 9 //int

	var k *int = &i
	var a **int = &k
	var l ***int = &a

	fmt.Println(***l)
	fmt.Println(**zero(a))
}

func zero(d **int) **int {
	var z int = 6
	var b *int = &z

	return &b
}
