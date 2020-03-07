package main

import "fmt"

func main() {

	//var tab = [5]int{1, 2, 3, 4, 5}

	//v := append(tab[:2], tab[3:]...)
	var tab = []int{2, 3, 5, 6, 1}
	var b = 2

	fmt.Println(sup(tab, &b))

}

func sup(tab []int, i *int) []int {

	tab2 := append(tab[:*i], tab[*i+1:]...)

	return tab2

}
