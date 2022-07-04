package main

import (
	"fmt"
)

// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

func main(){
	var t interface{}

	t = []int{1}
	checkType(t)

	t = "hello"
	checkType(t)

	t = make(chan int)
	fmt.Println("type:", typeFMT(t))

}

func checkType(t interface{}){

	switch v := t.(type) {
	case int:
		fmt.Println("int: ", v)
	case string:
		fmt.Println("string:", v)
	case bool:
		fmt.Println("bool:", v)
	case chan int:
		fmt.Println("chan int:", v)
	default:
		fmt.Printf("unknown: %T\n", v)
	}
}



func typeFMT(a interface{}) string {
	result:= fmt.Sprintf("%T", a)

	return result
}
