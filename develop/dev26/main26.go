package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func main() {
	s := "abcd"
	fmt.Println("check1:", chekStr(s))

	s1 := "abCdefAaf"
	fmt.Println("check1:", chekStr(s1))

	s2 := "аабцд"
	fmt.Println("check1:", chekStr(s2))

}
func tlower(s string) string {
	s = strings.ToLower(s)
	return s
}

func chekStr(s string) bool {
	lowerS := tlower(s)

	for i, _ := range lowerS {
		for j := i + 1; j < len([]rune(lowerS)); j++ {
			if string([]rune(lowerS)[i]) == string([]rune(lowerS)[j]) {
				return false
			}
		}
	}
	return true
}
