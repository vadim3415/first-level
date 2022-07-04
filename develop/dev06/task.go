package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.
func main() {
	fmt.Println("context")
	chWriterReader := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go writer(ctx, chWriterReader)

	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		cancel()
	}()
	for data := range chWriterReader {
		time.Sleep(time.Millisecond * 30)
		fmt.Println(data)
	}

	fmt.Println("printChannelTimer")
	time.Sleep(time.Second)
	channelTimer()

	fmt.Println("printChannelSignal")
	time.Sleep(time.Second)
	channelSignal()

	time.Sleep(time.Second)
	fmt.Println("Активных горутин", runtime.NumGoroutine())
}

//********context************

func writer(ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			close(ch)
			fmt.Println("поступила команда завершения работы")
			return
		default:
			rand.Seed(time.Now().UnixNano())
			time.Sleep(1 * time.Second)
			random := rand.Intn(10000)
			ch <- random
		}
	}
}

//********time************
func channelTimer() {

	newtimer := time.NewTimer(3 * time.Second)
	ch := make(chan int)

	for i := 0; i < 1000000; i++ {
		select {
		case <-newtimer.C:
			close(ch)
			fmt.Println("поступила команда завершения работы")
			return
		default:
			go printChannelTimer(ch)
			ch <- i
		}
	}
}

func printChannelTimer(ch chan int) {
	time.Sleep(1 * time.Second)
	s := <-ch
	fmt.Println(s)
}

//************ os.Signal AND context **************

func channelSignal() {
	chanSqr := make(chan float64)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	go shutdown(cancel)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go square(ctx, wg, chanSqr)
	}

	go func() {
		for v := range chanSqr {
			fmt.Println(v)
		}
	}()

	wg.Wait()
	close(chanSqr)
	time.Sleep(time.Second * 1)
}

func square(ctx context.Context, wg *sync.WaitGroup, chanSqr chan float64) {
	var i float64 = 10

	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
			time.Sleep(1 * time.Second)
			chanSqr <- math.Pow(i, 2)
			i++
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
