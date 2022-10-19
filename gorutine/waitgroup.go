package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//writeWaitGroup()
	writeWithMutex()

}

func writeWaitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(s int) {
			println(s)
			wg.Done()
		}(i)
	}

	wg.Wait()

	println("Exit")
}

func writeWithMutex() {
	var wg sync.WaitGroup
	var mx sync.Mutex

	count := 0
	start := time.Now()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(s int) {
			defer wg.Done()

			time.Sleep(time.Nanosecond)

			mx.Lock()
			count++
			mx.Unlock()

		}(i)
	}

	wg.Wait()

	delta_time := time.Now().Sub(start).Seconds()

	fmt.Printf("Count %d\n", count)
	fmt.Println(delta_time)
}
