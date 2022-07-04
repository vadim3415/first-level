package main

import (
	"fmt"
	"sort"
)

// Реализовать бинарный(двоичный) поиск встроенными методами языка.

func main() {
	arr := []int{1, 2, 3}
	x := 3
	// поиск производится в упорядоченном массиве ( по убыванию: a[i] <= x , по возрастанию  a[i] >= x)
	i := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= x
	})
	if i < len(arr) && arr[i] == x {
		fmt.Printf("Найдено %d по индексу %d в %v.\n", x, i, arr)
		// output: Найдено 3 по индексу 2 в [1 2 3].
	} else {
		fmt.Printf("Не найдено %d в %v.\n", x, arr)
	}
}
