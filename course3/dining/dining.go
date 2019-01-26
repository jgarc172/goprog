// Package main implements a solution to the Dining Philosophers problem
// by using a "host" that intermediates the timing for eating and allows
// only two diners at a time by using a buffered channel of two Philosophers
package main

import (
	"fmt"
	"runtime"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	id              int
	leftCS, rightCS *ChopS
}

func (p Philo) eat() {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.id)
		runtime.Gosched()
		fmt.Printf("finished eating %d\n", p.id)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
	done <- true
}

type Host struct {
}

func (h Host) canEat(p *Philo) {
	canEat <- p
}
func (h Host) serve() {
	fullDiners := 0
	for {
		select {
		case p := <-canEat:
			go p.eat()
		case <-done:
			fullDiners++
			if fullDiners == 5 {
				wg.Done()
				return
			}
		}
	}
}

var wg sync.WaitGroup
var done = make(chan bool)
var canEat = make(chan *Philo, 2)

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)

	host := new(Host)
	wg.Add(1)
	go host.serve()

	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%5]}
		go host.canEat(philos[i])
	}

	wg.Wait()
}
