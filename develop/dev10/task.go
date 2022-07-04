package main

import "fmt"

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
// 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
// градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func main() {
	s := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	result1 := groupTemp1(s)
	fmt.Println("result1", result1)

	result2 := groupTemp2(s)
	fmt.Println("result2", result2)
}

func groupTemp1(s []float64) map[int][]float64 {
	groupTemp := make(map[int][]float64)

	for _, v := range s {
		if v < -19.99 && v > -30 {
			groupTemp[-20] = append(groupTemp[-20], v)
		}
		if v > 9.99 && v < 20 {
			groupTemp[10] = append(groupTemp[10], v)
		}
		if v > 19.99 && v < 30 {
			groupTemp[20] = append(groupTemp[20], v)
		}
		if v > 29.99 && v < 40 {
			groupTemp[30] = append(groupTemp[30], v)
		}
	}

	return groupTemp
}

func groupTemp2(s []float64) map[int][]float64 {
	var groupTemp = make(map[int][]float64)

	for _, val := range s {
		group := int(val) / 10 * 10 //input 25.4: output: int(25.4) = 25, 25/10 = 2, 2*10 = 20
		groupTemp[group] = append(groupTemp[group], val)
	}

	return groupTemp
}
