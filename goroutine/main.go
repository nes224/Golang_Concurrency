package main

import (
	"fmt"
)

/*
func main() {
	start := time.Now()
	defer func(){
		fmt.Println(time.Since(start))
	}()

		// We will need to somehow spawn multiple processes or multiple go routines in order to perform these attacks concurrently.
		// hypothesis (N) = สมมุติฐาน
	evilNinjas := []string{"Tommy", "Johnny", "Bobby", "Andy"}
	for _, evevilNinja := range evilNinjas{
		go attack(evevilNinja)
	}
	// go attack(evevilNinja)  go -> go routines
	time.Sleep(time.Second*2)
}

func attack(target string) {
	fmt.Println("Throwing naja stars at", target)
	time.Sleep(time.Second)
}
*/

// how we can spawn multiple processes or go routines
// in order to do perform nultiple tasks at once.

/*
func main(){
	now := time.Now()
	defer func(){
		fmt.Println(time.Since(now))
	}()
	// If you want to create a slice you should use the built-in make function:
	smokeSignal := make(chan bool)
	evilNinja := "Tommy"
	go attack(evilNinja,smokeSignal)
	fmt.Println(<-smokeSignal)

}

func attack(target string, attacked chan bool){
	time.Sleep(time.Second)
	fmt.Println("Throwing ninja stars at",target)
	attacked <- true // ch <- v    // Send v to channel ch.
}
*/
// buffered channels

func main() {
	channel := make(chan string, 2)
	// fatal error: all groutines are asleep - deadlock!
	channel <- "First Message"
	channel <- "Second Message"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

