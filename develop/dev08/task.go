package main

import (
	"fmt"
)

func checkBit(num int64, pos uint) bool {
	val := num & (1 << pos)
	return val > 0
}

func ChangeNBit(num int64, pos uint) int64 {

	pos--
	// установить бит в ноль
	if checkBit(num, pos) {
		num &= ^(1 << pos)
		// установить бит в 1
	} else {
		num |= (1 << pos)
	}
	return num
}

func main() {
	num := int64(64)
	fmt.Printf("число: %d\n : в двоичной %b\n", num, num)

	changedNum := ChangeNBit(num, 3)
	fmt.Printf("число: %d\n : в двоичной %b\n", changedNum, changedNum)
}