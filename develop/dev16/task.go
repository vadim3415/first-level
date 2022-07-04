package main

import (
	"fmt"
	"sort"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами
// языка.

func main(){
	m:= [5]int{1,5,2,4,3}
	sorted(m)
}

func sorted (m [5]int){
	sort.Ints(m[:])
 	fmt.Println(m)
}