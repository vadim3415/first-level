package main

import "fmt"

// Удалить i-ый элемент из слайса.

func main() {
	s := []int{1, 2, 3, 4, 5}
	index := 1

	result := deletIndex(s, index)
	fmt.Println(result) // output: [1 3 4 5]

	s2 := []int{1, 2, 3, 4, 5}
	result2 := deletIndex2(s2, index)
	fmt.Println(result2) // output: [1 5 3 4]
}

// удаление элемента без изменения порядка
func deletIndex(arr []int, index int) []int {
	arr = append(arr[:index], arr[index+1:]...)

	return arr
}

// удаление элемента с изменением порядка
func deletIndex2(arr []int, index int) []int {
	arr[index] = arr[len(arr)-1] //ставим последний элемент на место удаляемого [1 2 3 4 5], вместо 2 будет 5
	arr = arr[:len(arr)-1]       //обрезаем слайс(копия без последнего элемента)

	return arr
}
