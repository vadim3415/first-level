package main

import (
	"fmt"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.

func main(){
	newtimer := time.NewTimer(5 * time.Second)
	ch:= make(chan int)
	
	for i:= 0; i<1000000; i++{
	select {
	case <-newtimer.C:
		close(ch)
		fmt.Println("exit")
		return
	default:
			time.Sleep(1 * time.Second)
			go print(ch)
			ch <- i
		}
	}
}

func print (ch chan int){
	s := <- ch
	fmt.Println(s)
}