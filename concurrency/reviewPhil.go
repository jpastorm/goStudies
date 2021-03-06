package main

import (
	"fmt"
	"sync"
	"time"
)

var numberOfPhilos int = 5

type Permission struct {
	v  int
	mu sync.Mutex
	wg sync.WaitGroup
}

func (per *Permission) Inc() {
	per.mu.Lock()
	defer per.mu.Unlock()
	per.v++
	if per.v >= numberOfPhilos {
		per.mu.Unlock()
		per.wg.Wait()
		per.wg.Add(1)
		return
	}
	if per.v == numberOfPhilos-1 {
		per.wg.Add(1)
		return
	}
}

func (per *Permission) Dec() {
	per.mu.Lock()
	per.v--
	if per.v == numberOfPhilos-1 {
		per.wg.Done()
		return
	}
	if per.v == numberOfPhilos-2 {
		per.wg.Done()
	}
	per.mu.Unlock()
}

var globalPer Permission
var globalWg sync.WaitGroup

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
	v               int
}

func (p Philo) eat() {
	for i := 0; i < 3; i++ {
		globalPer.Inc()
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.v+1)
		time.Sleep(1000)
		fmt.Printf("finishing eating %d\n", p.v+1)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		globalPer.Dec()
	}
	globalWg.Done()
}

func main() {
	// Initialization
	CSticks := make([]*ChopS, numberOfPhilos)
	for i := 0; i < numberOfPhilos; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, numberOfPhilos)
	for i := 0; i < numberOfPhilos; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%numberOfPhilos], i}
	}
	// Start the Dining
	for i := 0; i < 5; i++ {
		globalWg.Add(1)
		go philos[i].eat()
	}
	globalWg.Wait()
	fmt.Println("done.")
}
