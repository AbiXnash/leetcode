package main

// TODO: this logic is totally wrong i need to implement stack for this

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

	for i := 0; i < len; i += 2 {
		curr := string(s[i])
		if expected := parenthesis[curr]; expected != getNext(s, i) {
			return false
		}
	}

	return true
}

func getNext(s string, index int) string {
	if index+1 < len(s) {
		return string(s[index+1])
	}
	return ""
}
