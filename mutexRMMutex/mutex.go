package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock   sync.Mutex
	rwLock sync.RWMutex
	count  int
)

func main() {
	readAndWrite()
}

func basics() {
	iteration := 1000
	for i := 0; i < iteration; i++ {
		go increment()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Resulted count is:", count)
}

func readAndWrite(){
	go read()
	go read()
	go write()

	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}

func read(){
	rwLock.RLock()
	defer rwLock.RUnlock()

	fmt.Println("Read locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Reading unlocking")
}

func write(){
	rwLock.Lock()
	defer rwLock.Unlock()

	fmt.Println("write locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Write unlocking")
}

func increment() {
	lock.Lock()
	count++
	lock.Unlock()
}
