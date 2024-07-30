package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

const numberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= numberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Println("Received order", pizzaNumber)
		rnd := rand.Intn(12) + 1
		msg := ""
		succes := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		fmt.Println("Making pizza", pizzaNumber, ". It will take", delay, "second.")

		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("We ran out ingredients for pizza %d", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("The cook quit while making pizza %d", pizzaNumber)
		} else {
			succes = true
			msg = fmt.Sprintf("Pizza order %d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: numberOfPizzas,
			message:     msg,
			success:     succes,
		}
		return &p
	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	var i = 0
	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
		}
		i = currentPizza.pizzaNumber
		select {
		case pizzaMaker.data <- *currentPizza:
		case quitChan := <-pizzaMaker.quit:
			close(pizzaMaker.data)
			close(quitChan)
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	color.Cyan("The pizzeria is open for business!")
	color.Cyan("----------------------------------")

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	go pizzeria(pizzaJob)

	for i := range pizzaJob.data {
		if i.pizzaNumber <= numberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order %d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizzas ...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("Error closing channel!", err)
			}
		}
	}

	color.Cyan("-----------------")
	color.Cyan("Done for the day.")
}

func test(err chan error) error {
	if err != nil {
		return <-err
	}
	return nil
}
