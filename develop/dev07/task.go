package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

type Storage struct {
	data map[int]string
	mtx  sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[int]string),
		mtx:  sync.RWMutex{},
	}
}

func main() {

	s := NewStorage()
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go s.Read(&wg)
		go s.Write(&wg)
	}
	wg.Wait()
}

func (s *Storage) Read(wg *sync.WaitGroup) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Println(s.data)
	}
}

func (s *Storage) Write(wg *sync.WaitGroup) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	defer wg.Done()

	for i := 0; i < 10; i++ {
		s.data[i] = fmt.Sprintf("msg = %d", i)
	}
}
