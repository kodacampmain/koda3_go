package filereader

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Read(filepath string) error {
	log.Println("Opening file...")
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer func() {
		log.Println("Closing file...")
		file.Close()
	}()
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			log.Println("Continue...")
		}
	}()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		panic("filepath is not a file")
		// return nil
	}

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		return scanner.Err()
	}

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}
