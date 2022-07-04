package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
// собственное множество.

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	many(arr)
}

func many(arr []string) {
	m := make(map[string]struct{})
	var output []string

	for _, v := range arr {
		m[v] = struct{}{}
	}

	for key, _ := range m {
		output = append(output, key)
	}

	fmt.Println(output)
}
