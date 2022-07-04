package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры
// Human (аналог наследования).

type Printer interface{
	print()
	printAge() int
}

type Human struct{
Name string
Age int
}
//встравивание(наследует  все методы)
type Action struct{
	Human
}
//композиция
type Action2 struct{
	Human Human
}

func (h Human) print(){
	fmt.Println("Human method:", h.Name, h.Age, "struct:", h)
}

func (a Action2) print(){
	fmt.Println("Action method:", a.Human.Name, a.Human.Age, "struct:", a)
}

func (a Action2) printAge() int{
	return a.Human.Age
}


func main(){
hum:= Human{Name: "Bob", Age: 31}
act:= Action{Human{Name: "Alice", Age: 25} }
act2:= Action2{Human{Name: "Alice", Age: 25} }

hum.print() // Human method:
act.print() // Human method:

act2.print() // Action method:
act2.Human.print() // Human method:

age(act2)

}

func age(p Printer){
	fmt.Println(p.printAge())
}
