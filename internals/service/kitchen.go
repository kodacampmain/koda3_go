package service

import (
	"fmt"
	"sync"
	"time"
)

func Kitchen(foods, drinks chan string, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case food := <-foods:
			fmt.Printf("Chef mulai memasak %s\n", food)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("Chef selesai memasak %s\n", food)
		case drink := <-drinks:
			fmt.Printf("Barista mulai membuat minuman %s\n", drink)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("Barista selesai membuat minuman %s\n", drink)
		case <-done:
			fmt.Println("Kitchen closing...")
			return
		}
	}
}
