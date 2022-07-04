package main

import (
	"fmt"
	"time"
)


func MyTimer(t time.Duration) <-chan time.Time{
	times:= time.After(t)
	return times
}

func MySleep(t time.Duration){
	<-time.After(t)
}

func main() {
	start := time.Now()
  channelTimer()
 
  fmt.Println("время выполнения программы: ", time.Since(start))
}

func channelTimer() {
	newtimer := MyTimer(time.Second * 4)
	ch := make(chan int)

	for i := 0; i < 1000000; i++ {
		select {
		case <-newtimer:
			close(ch)
			fmt.Println("exit")
			return
		default:
			go printChannelTimer(ch)
			ch <- i
		}
	}
}

func printChannelTimer(ch chan int) {
	MySleep(time.Second * 1)
	s := <-ch
	fmt.Println(s)
}