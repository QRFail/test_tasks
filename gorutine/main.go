package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())

	go showNumners()

	fmt.Println(runtime.NumCPU())

	//time.Sleep(time.Second)

	fmt.Println("Exit")
}

func showNumners() {
	var i int = 0
	for i < 100 {
		i++
		fmt.Printf("%d\n", i)
	}
}
