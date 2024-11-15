package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Recipe struct {
	Name    string
	Minutes int
}

func (r Recipe) String() string {
	return fmt.Sprintf("%v: (%d minutes)", r.Name, r.Minutes)
}

var recipes = []Recipe{
	{
		Name:    "Spaghetti Carbonara",
		Minutes: 20,
	},
	{
		Name:    "Spaghetti Bolognese",
		Minutes: 30,
	},
	{
		Name:    "Spaghetti al Pesto",
		Minutes: 15,
	},
	{
		Name:    "Spaghetti Aglio e Olio",
		Minutes: 10,
	},
}

func main() {
	const generators = 10
	orders := make(chan Recipe)
	wg := sync.WaitGroup{}
	for i := 0; i < generators; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			randomIndex := rand.Intn(len(recipes))
			recipe := recipes[randomIndex]
			orders <- recipe
		}()
	}
	go func() {
		wg.Wait()
		close(orders)
	}()

	for order := range orders {
		fmt.Println(order)
	}
}
