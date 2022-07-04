package main

import (
	"fmt"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(2*2+4*2+6*2....) с использованием конкурентных вычислений.

func main() {

	chConv := make(chan int)  // канал для отправки и чтения значений из среза.
	chSqare := make(chan int) // канал для отправки и чтения значений возведнных в квадрат
	quit := make(chan int)    // канал для остановки

	m := []int{2, 4, 6, 8, 10}

	go conv(m, chConv)
	go square(chConv, chSqare)
	go sumSquare(chSqare, quit)
	for {
		q := <-quit
		if q == 0 {
			break
		}
	}
	fmt.Printf("sum square: %d = %d\n", m, sumSqu)
}

var sumSqu int

func conv(m []int, chConv chan int) {
	for i := 0; i < len(m); i++ {
		chConv <- m[i]
	}
	close(chConv)
}

func square(chConv chan int, chSquare chan int) {
	for {
		number := <-chConv
		if number == 0 {
			break
		}
		chSquare <- number * number
	}
	close(chSquare)
}

func sumSquare(chSquare chan int, quit chan int) {
	for {
		sq := <-chSquare
		if sq == 0 {
			break
		}
		sumSqu = sumSqu + sq
		quit <- sumSqu
	}
	close(quit)
}
