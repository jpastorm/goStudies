package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
}

var wg sync.WaitGroup
var philospherCounter int = 0

func (p Philo) eat(index int) {
	i := 0
	for i < 3 {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat ", index)
		fmt.Println("finishing eating ", index)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		i++

	}
	wg.Done()
}

func main() {

	Csticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		Csticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		// if i < (i+1)%5 {
		philos[i] = &Philo{Csticks[i], Csticks[(i+1)%5]}
		// } else {
		// 	philos[i] = &Philo{Csticks[(i+1)%5], Csticks[i]}
		// }

	}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(i)
	}
	wg.Wait()

}
