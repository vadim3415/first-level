package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	str := "главрыба"
	fmt.Printf("reverse(%s): [%s]\n", str, reverse(str))

	str2 := "abcd"
	fmt.Printf("reverse(%s): [%s]\n", str2, reverse(str2))

	str3 := "日本語"
	fmt.Printf("reverse(%s): [%s]\n", str3, reverse(str3))
}

func reverse(s string) string {
	lenS := len([]rune(s))
	m := make([]rune, lenS)

	for _, c := range s {
		lenS--
		m[lenS] = c
	}
	return string(m)
}
