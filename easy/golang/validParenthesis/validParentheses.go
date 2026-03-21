package main

import (
	"fmt"
)

var parenthesis = map[string]string{
	"{": "}",
	"[": "]",
	"(": ")",
}

func main() {
	result := IsValid("(}){")
	fmt.Println(result)
}

func IsValid(s string) bool {
	len := len(s)
	if len%2 != 0 {
		return false
	}

	return true
}
