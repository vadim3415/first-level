package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.

//Программа должна завершаться по нажатию Ctrl+C.
//Выбрать и обосновать способ завершения работы всех воркеров.

func main() {
	workersNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf(err.Error())
	}

	chWriteRead := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go shutdown(cancel)

	for i := 0; i < workersNum; i++ {
		go printer(ctx, chWriteRead, i)
	}
	
	conv(ctx, chWriteRead)
	
	time.Sleep(time.Second * 1)
	fmt.Println("Активных горутин", runtime.NumGoroutine())
}

func printer(ctx context.Context, chRead chan int, i int) {
	for {
		select{
		case <-ctx.Done():
			return
		default:
			for v := range chRead {
				fmt.Println(v, "chan", i)
			}
		}
	}
}

func conv(ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			rand.Seed(time.Now().UnixNano())
			time.Sleep(1 * time.Second)
			random := rand.Intn(10000)
			ch <- random
		}
	}
}

func shutdown(cancel context.CancelFunc) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("поступила команда завершения работы") 
	
	cancel() 
}

