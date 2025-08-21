package utils

import (
	"log"
	"sync"
	"time"
)

func Task(taskId int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 3 {
		log.Printf("Task %d: step %d", taskId, i)
		time.Sleep(100 * time.Millisecond)
	}
}

var counter int

func Read(wg *sync.WaitGroup, channel chan string, mutex *sync.RWMutex, unlock bool) {
	defer wg.Done()
	defer func() {
		if unlock {
			mutex.RUnlock()
		}
	}()
	log.Println(counter)
	// channel <- strconv.Itoa(counter)
}

func Write(wg *sync.WaitGroup, strChan chan string, mutex *sync.RWMutex) {
	defer wg.Done()
	defer mutex.Unlock()
	counter++
	// log.Println("Writing Counter")
	// strChan <- "Update Counter"
}
