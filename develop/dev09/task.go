package main

import (
	"fmt"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.

func main() {

	chConv := make(chan int) // канал для чисел из массива
	chMult := make(chan int) // канал для операций x*2

	m := []int{2, 4, 6, 8, 10}

	go conv(m, chConv)
	go multiplication(chConv, chMult)
	var sumSq int

	for {
		sumSq = <-chMult

		if sumSq == 0 {
			break
		}
		fmt.Println(sumSq)
	}
}

func conv(m []int, chConv chan int) {
	for i := 0; i < len(m); i++ {
		chConv <- m[i]
	}
	close(chConv)
}

func multiplication(chConv chan int, chSquare chan int) {
	for {
		mult := <-chConv
		if mult == 0 {
			break
		}
		chSquare <- mult * 2
	}
	close(chSquare)
}
