package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hardTask(name string) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Hard Task %s...\n", name)
	}
	fmt.Printf("Hard Task %s DONE\n", name)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hardTask(strconv.Itoa(i))
	}

	wg.Wait()

	fmt.Printf("DONE")
}
