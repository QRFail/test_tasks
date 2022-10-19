package main

import "sync"

func main() {

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
