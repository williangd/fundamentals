package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "An implicity typed string"
	fmt.Printf("str1: %v:%T\n", str1, str1)
	var str2 string = "An explicity typed string"
	fmt.Printf("str2: %v:%T\n", str2, str2)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str1))

	lValue := "hello"
	uValue := "HELLO"
	fmt.Println("Equal?", (lValue == uValue))
	fmt.Println("Equal Non-case?", strings.EqualFold(lValue, uValue))

	fmt.Println("Contains exp?", strings.Contains(str1, "exp"))
	fmt.Println("Contains exp?", strings.Contains(str2, "exp"))
}
