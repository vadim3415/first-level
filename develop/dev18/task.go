package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

type counter struct {
	count int32
}

func main() {
	count := &counter{}
	newTimer := time.NewTimer(1 * time.Second)

	for {
		select {
		case <-newTimer.C:
			fmt.Println("end: ", count.count)
			return
		default:
			go atomCount(count)
		}
	}
}

func atomCount(con *counter) {
	atomic.AddInt32(&con.count, 1)
}
