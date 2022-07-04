package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func main() {

	arr1 := []int{4, 9, 5, 9, 7}
	arr2 := []int{1, 9, 4, 9, 8, 4, 3, 9}
	intersect(arr1, arr2)

}

func intersect(arr1, arr2 []int) {
	m := make(map[int]int)
	var output []int

	// в ключ карты записываем число из 1-го среза, в значение кол-во повторений
	for _, v := range arr1 {
		m[v] += 1
	}
	// проверям число из 2-го слайса на наличие в ключе карты и кол-во повторений числа
	for _, v := range arr2 {
		if _, ok := m[v]; ok && m[v] > 0 {
			output = append(output, v)
			m[v] -= 1
		}

	}
	fmt.Println(output)
}
