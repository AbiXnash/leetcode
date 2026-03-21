package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "({})"
	fmt.Println(a)
	result := isValid(a)
	fmt.Println(result)
}

var parenthesis = map[string]string{
	"{": "}",
	"[": "]",
	"(": ")",
	"}": "{",
	"]": "[",
	")": "(",
}

func isValid(s string) bool {
	len := len(s)

	if len%2 != 0 {
		return false
	}

	for i := 0; i < len/2; i++ {
		return strings.Compare(string(s[i]), parenthesis[string(s[len-i-1])]) == 0
	}

	return true
}
