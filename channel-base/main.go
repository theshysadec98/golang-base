package main

import (
	"fmt"
	"strings"
	"time"
)

func shout(ping chan string, fps chan string) {
	for {
		s, ok := <-ping
		if ok {
			fmt.Println(s)
		}
		fps <- fmt.Sprintf("%s", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	fps := make(chan string)
	go shout(ping, fps)

	time.Sleep(3 * time.Second)
	fmt.Println("Type something and press ENTER (enter Q to quit)")

	for {
		fmt.Print("-> ")
		var input string
		_, _ = fmt.Scan(&input)

		if input == "Q" {
			break
		}

		ping <- input
		response := <-fps

		fmt.Println("Response:", response)
	}

	fmt.Println("Close process.")
	close(ping)
	close(fps)
}
