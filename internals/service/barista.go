package service

import (
	"fmt"
	"sync"
	"time"
)

func Barista(orders chan string, baristaId int, wg *sync.WaitGroup) {
	defer wg.Done()
	// pesanan akan datang melalui channel
	for order := range orders {
		fmt.Printf("Barista %d making orders of %s\n", baristaId, order)
		time.Sleep(2 * time.Second)
		fmt.Printf("Barista %d finished orders of %s\n", baristaId, order)
	}
	fmt.Println("No more orders. Barista closing shop...")
}
