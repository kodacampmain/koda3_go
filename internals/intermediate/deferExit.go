package intermediate

import (
	"fmt"
	"log"
	"os"
)

func DeferredFunction() {
	defer fmt.Println("Selesai")
	fmt.Println("Dimulai")
}

func Exit(code int, err error) {
	log.Println(err.Error())
	os.Exit(code)
}
