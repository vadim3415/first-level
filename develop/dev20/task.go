package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func main() {
	str := "snow dog sun cat rain"
	fmt.Printf("reverse(%s): [%s]\n", str, reverse(str))
	str2 := "привет мир"
	fmt.Printf("reverse(%s): [%s]\n", str2, reverse(str2))
}

func reverse(s string) string {

	str := strings.Split(s, " ")
	lenS := len(str)
	for j := 1; j < lenS/2+1; j++ {
		str[j-1], str[lenS-j] = str[lenS-j], str[j-1]
	}
	joinStr := strings.Join(str, " ")

	return joinStr
}
