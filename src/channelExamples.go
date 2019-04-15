package main

import (
	"fmt"
	"runtime"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < 200; i++ {
		// _ = <-time.After(100 * time.Millisecond)

		c <- i
		fmt.Println("********************put in channel", i)
		x, y = y, x+y
	}
	fmt.Println("Close", len(c))
	close(c)
}

func examClosechannel() {
	c := make(chan int, 20)
	go fibonacci(cap(c), c)
	fmt.Println("Begin", len(c))
	var readOnlyChan <-chan int
	readOnlyChan = c
	for va := range readOnlyChan {
		// fmt.Println("Going", len(c), cap(c), va)
		fmt.Println("___________________read in channel ", va)

		_ = va

	}
	fmt.Println("After", len(c))
	for _, v := range [10]int{1, 23, 3, 3, 3, 3, 3, 3} {
		// println("**********", v)
		_ = v
	}
}

func examSelect() {
	runtime.GOMAXPROCS(1)
	intChannel := make(chan int, 1)
	stringChannel := make(chan string, 1)
	intChannel <- 1
	stringChannel <- "hello"
	select {
	case value := <-intChannel:
		fmt.Println(value)
	case value := <-stringChannel:
		panic(value)
	}
	fmt.Println(123)
}
