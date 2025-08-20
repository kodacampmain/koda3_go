package intermediate

import (
	"fmt"
	"log"
)

func MustDo(isError bool) {
	defer func() {
		fmt.Println("cleanup process")
	}()
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("panic recovered")
			log.Printf("panic: %v", r)
		}
	}()
	fmt.Println("processing data 1")
	if isError {
		panic("unexpected error")
	}
	fmt.Println("processing data 2")
}
