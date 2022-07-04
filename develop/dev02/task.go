package main

import (
	"fmt"
	"math"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	m := [5]float64{2, 4, 6, 8, 10}
	ch := make(chan float64)

	for i := 0; i < len(m); i++ {
		go squares(ch, m[i])
		fmt.Println(<-ch)
	}
	close(ch)
	println("exit")
}

func squares(ch chan float64, i float64) {
	ch <- math.Pow(i, 2)
}
