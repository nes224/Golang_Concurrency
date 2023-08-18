package main

import (
	"fmt"
	"sync"
	// "sync"
	"math/rand"
	"time"
)


func main(){
	gettingReadyForMission()
}

var ready bool

func gettingReadyForMissionWithCond(){
	cond := sync.Cond(&sync.Mutex{})
	go gettingReadyWithCond(& cond)
	workIntervals := 0
	for !ready{
		time.Sleep(5 * time.Second)
		workIntervals++
	}

	fmt.Printf("We are now ready! after %d work intervals. \n",workIntervals)
}

func gettingReadyForMission(){
	go gettingReady()
	workIntervals := 0
	for !ready{
		time.Sleep(5 * time.Second)
		workIntervals++
	}

	fmt.Printf("We are now ready! after %d work intervals. \n",workIntervals)
}

func gettingReadyWithCond(cond *sync.Cond){
	sleep()
	ready = true
	cond.Signal()
}

func gettingReady(){
	sleep()
	ready = true
}

func sleep(){
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1 + rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}

func NewCond(l Locker) *Cond{
	return &Cond{L:l}
}

/* 

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	// This function represents a producer that signals the consumer when data is available.
	producer := func(id int) {
		defer wg.Done()
		time.Sleep(time.Duration(id) * time.Second) // Simulate producing data
		fmt.Printf("Producer %d: Data is ready\n", id)

		mu.Lock()
		cond.Signal() // Signal the waiting consumer
		mu.Unlock()
	}

	// This function represents a consumer that waits for data from the producer.
	consumer := func(id int) {
		defer wg.Done()
		fmt.Printf("Consumer %d: Waiting for data...\n", id)

		mu.Lock()
		cond.Wait() // Wait for data signal
		mu.Unlock()

		fmt.Printf("Consumer %d: Got the data!\n", id)
	}

	wg.Add(3)
	go producer(1)
	go consumer(1)
	go consumer(2)
	wg.Wait()
}
*/