package main

import (
	"fmt"
	"os"
)

var romanValues = map[string]int{
	"":  0,
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("Test [%v]: [%d]\n", arg, RomanToInt(arg))
	}
}

func RomanToInt(s string) int {
	var theNumber int
	len := len(s)
	for i := range s {
		curr := romanValues[string(s[i])]
		next := romanValues[getNext(s, len, i)]

		if curr < next {
			theNumber -= curr
		} else {
			theNumber += curr
		}
	}

	return theNumber
}

func getNext(s string, len, i int) string {
	if i+1 < len {
		return string(s[i+1])
	} else {
		return ""
	}
}
