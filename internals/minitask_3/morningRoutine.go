package minitask3

import (
	"fmt"
	"sync"
	"time"
)

func DoTask(taskname string, durationInMs int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Memulai pekerjaan %s\n", taskname)
	time.Sleep(time.Duration(durationInMs) * time.Millisecond)
	fmt.Printf("Selesai mengerjakan %s\n", taskname)
}
