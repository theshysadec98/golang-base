package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

const seatingCapacity = 2
const arrivalRate = 100
const cutDuration = 100 * time.Millisecond
const timeOpen = 10 * time.Second

func main() {
	// seed our random number generator
	rand.Seed(time.Now().UnixNano())

	//Print message welcome
	color.Yellow("The sleeping Barber Problem")
	color.Yellow("---------------------------")

	//Create channel if we need any
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	//create the barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarber:  0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The shop is open for the day.")

	// add barber
	shop.addBarber("Trundle")

	//start the barbershop as a goroutine
	shopClose := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClose <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// add clients
	i := 1

	go func() {
		for {
			randomMilliseconds := rand.Int() % (2 * arrivalRate)

			select {
			case <-shopClose:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.addBarber(fmt.Sprintf("Client %d", i))
				i++
			}
		}
	}()
	//block until the barbershop is closed
	time.Sleep(5 * time.Second)
}
