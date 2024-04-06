package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myChannel := make(chan int)
	wg := &sync.WaitGroup{}

	//Adding <-chan makes the goroutine Read only
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myChannel

		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myChannel, wg)

	//Adding chan<- makes the goroutine send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChannel <- 5
		close(myChannel)
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
