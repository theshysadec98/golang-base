package main

import (
	"fmt"
	"time"
)

func listenChanel(ch chan int) {
	for {
		i := <-ch
		fmt.Println("Got", i, "from channel")
		time.Sleep(1 * time.Second)
	}
}
func main() {
	ch := make(chan int, 10)

	go listenChanel(ch)

	for i := 0; i <= 100; i++ {
		fmt.Println("seeding", i, "to channel ...")
		ch <- i
		fmt.Println("send", i, "to channel")
	}
	fmt.Println("Done!")

	close(ch)
}
