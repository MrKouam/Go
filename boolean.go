package main

import (
	"os"

	"github.com/01-edu/z01"
)


func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {

	if nbr %2 == 0 {
		return true
	} else {
		return false 
	}
}

func main() {

	var EvenMsg  = "I have an even number of arguments"
	var OddMsg   = "I have an odd number of arguments"	
	//lengthOfArg = len(os.Args[])
	lengthOfArg := len(os.Args[1:])

	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
