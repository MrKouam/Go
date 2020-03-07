# Go


package main

import (
	"fmt"
)

func main() {
	var tab = [5]string{"201RichK", "Franck", "zulu", "Mathieu", "Anna"}
	i := 0

	for ; i < len(tab); i++ {
		fmt.Println(tab[i])

	}
}



package main

import (
	"fmt"
)

func main() {
	var tab = [5]string{"201RichK", "Franck", "zulu", "Mathieu", "Anna"}

	fmt.Println(tab[:4])
}
