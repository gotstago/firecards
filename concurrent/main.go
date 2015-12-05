package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var logs = map[time.Time]string{}
var mutex sync.Mutex

type order struct {
	dish     string
	num      int
	duration time.Duration
}

type chef struct {
	station int
	name    string
}

func newChef(name string, station int) *chef {
	return &chef{
		station: station,
		name:    name,
	}
}
func (c *chef) rest() {
	// minimal rest required by law
	fmt.Printf("\t%s is resting\n", c.name)
	//time.Sleep(100 * time.Millisecond)
	//rand.Int31n(1000)
	time.Sleep(time.Duration(rand.Int31n(10000)) * time.Millisecond)
}
func (c *chef) cook(o *order) {
	fmt.Printf("\t%s is cooking a %s (order %d)\n", c.name, o.dish, o.num)
	time.Sleep(o.duration)
	mutex.Lock()
	logs[time.Now()] = fmt.Sprintf("%s cooked a %s", c.name, o.dish)
	mutex.Unlock()
	c.rest()
}

func five(orders []*order, chefs []*chef) {
	wg := &sync.WaitGroup{}
	orderChan := make(chan *order)
	for _, c := range chefs {
		wg.Add(1)
		go func(c *chef) {
			for o := range orderChan {
				c.cook(o)
			}
			wg.Done()
		}(c)
	}
	for i, order := range orders {
		fmt.Printf("order %d: %s ", i+1, order.dish)
		orderChan <- order
	}
	close(orderChan)
	wg.Wait()
}

func main() {
	chefs := []*chef{newChef("François", 2), newChef("Rose", 3), newChef("Bill", 3)}
	orders := []*order{
		{"Blanquette de veau", 1, 1500 * time.Millisecond},
		{"Soupe à l'oignon", 2, 4850 * time.Millisecond},
		{"Blanquette de veau", 3, 2500 * time.Millisecond},
		{"Soupe à l'oignon", 4, 7850 * time.Millisecond},
		{"Blanquette de veau", 5, 1500 * time.Millisecond},
		{"Soupe à l'oignon", 6, 1850 * time.Millisecond},
		{"Soupe à l'oignon", 7, 1850 * time.Millisecond},
		{"Blanquette de veau", 8, 1500 * time.Millisecond},
		{"Soupe à l'oignon", 9, 4850 * time.Millisecond},
		{"Blanquette de veau", 10, 2500 * time.Millisecond},
		{"Soupe à l'oignon", 11, 7850 * time.Millisecond},
		{"Blanquette de veau", 12, 1500 * time.Millisecond},
		{"Soupe à l'oignon", 13, 1850 * time.Millisecond},
		// […]
	}
	startT := time.Now()
	five(orders, chefs)
	fmt.Printf("all done in %s, closing the kitchen\n", time.Since(startT))
	fmt.Println("logs:")

	for t, entry := range logs {
		fmt.Printf("%s: %s\n", t, entry)
	}
}
