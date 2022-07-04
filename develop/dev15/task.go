package main

import (
	"fmt"

	"log"
	"os"
)
// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.
// var justString string
// func someFunc() {
// v := createHugeString(1 << 10)
// justString = v[:100]
// }
// func main() {
// someFunc()
// }
func main() {
	someFunc()
}


func someFunc() {
	var justString string
	v := createHugeString(1 << 10)
	fmt.Println(cap([]rune(v)), len(v))

	justString = string([]rune(v)[:100])
	fmt.Println(justString,  cap([]rune(justString)), len(justString))

}

func createHugeString(size int) string {
	file, err := os.Open("develop/dev15/text.txt")
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, size)
	if _, err := file.Read(buf); err != nil {
		log.Fatal(err)
	}

	return string(buf)
}

