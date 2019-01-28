// Package main implements a solution to the Dining Philosophers problem
// by using a "host" that intermediates the timing for eating and allows
// only two diners at a time by using a buffered channel of two Philosophers
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)

	host := NewHost(philos, CSticks)
	host.Serve()
}

// ChopS is a representation of a single chopstick
type ChopS struct{ sync.Mutex }

// Philo is a representation of a Dining Philosopher
type Philo struct {
	id              int
	leftCS, rightCS *ChopS
	host            *Host
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
	p.host.Done()
}

// Host coordinates dinner for the Dining Philosophers
type Host struct {
	philos  []*Philo
	cSticks []*ChopS
	canEat  chan *Philo
	done    chan bool
}

// NewHost creates a new instance of Host
func NewHost(phils []*Philo, chs []*ChopS) *Host {
	canEat := make(chan *Philo, 2)
	var done = make(chan bool)

	return &Host{phils, chs, canEat, done}
}
func (h Host) waitOn(p *Philo) {
	h.canEat <- p
}

//Done is called by diners after finishing eating
func (h Host) Done() {
	h.done <- true
}

//Serve starts serving diners until finished
func (h Host) Serve() {

	for i := 0; i < 5; i++ {
		h.philos[i] = &Philo{i, h.cSticks[i], h.cSticks[(i+1)%5], &h}
		go h.waitOn(h.philos[i])
	}

	dinersServed := 0
	for {
		select {
		case p := <-h.canEat:
			go p.eat()
		case <-h.done:
			dinersServed++
			fmt.Println("\nhost finished serving a diner")
			fmt.Println()
			if dinersServed == 5 {
				return
			}
		}
	}
}
