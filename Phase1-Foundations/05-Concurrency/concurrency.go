package main

import (
	"fmt"
	"sync"
	"time"
)

type order struct {
	Details string
	Delay   int
}

func processOrder(o order, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("processing %s \n", o.Details)
	time.Sleep(time.Second * time.Duration(o.Delay))
	fmt.Printf("%s processed \n", o.Details)
}

func main() {
	orders := []order{
		{Details: "order 1", Delay: 1},
		{Details: "order 2", Delay: 4},
		{Details: "order 3", Delay: 3},
		{Details: "order 4", Delay: 6},
	}
	wg := sync.WaitGroup{}
	for _, order := range orders {
		wg.Add(1)
		go processOrder(order, &wg)
	}
	wg.Wait()
}
