package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.

// Mul - произведение x*y
func multiply(x, y *big.Int) *big.Int {
	return new(big.Int).Mul(x, y)
}

// Div деление x/y где y != 0 
func divide(x, y *big.Int) *big.Int {
	return new(big.Int).Div(x, y)
}

// Add - сумма x+y
func add(x, y *big.Int) *big.Int {
	return new(big.Int).Add(x, y)
}

// Sub вычетание x-y 
func sub(x, y *big.Int) *big.Int {
	return new(big.Int).Sub(x, y)
}

func main() {

	bigIntX := big.NewInt(int64(1 << 40)) // 1 099 511 627 776
	bigIntY := big.NewInt(int64(1 << 30)) // 1 073 741 824

	fmt.Println("bigIntA", bigIntX)
	fmt.Println("bigIntB", bigIntY)

	fmt.Println("умножение: ", multiply(bigIntX, bigIntY))

	fmt.Println("деление: ", divide(bigIntX, bigIntY))

	fmt.Println("сумма: ", add(bigIntX, bigIntY))

	fmt.Println("вычетание: ", sub(bigIntX, bigIntY))

}
