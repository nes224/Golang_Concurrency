package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var missionComplete bool 

func main(){
	var wg sync.WaitGroup
	var once sync.Once
	wg.Add(100)
	for i :=0 ; i< 100; i++{
		go func(){
			if foundTreasure(){
				once.Do(markMissionCompleted)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	checkMissionCompletion()

}

func checkMissionCompletion() {
	if missionComplete {
		fmt.Println("Mission is now completed.")
	} else {
		fmt.Println("Mission was failure.")
	}
}

func markMissionCompleted() {
	missionComplete = true 
}

func foundTreasure() bool  {
	rand.Seed(time.Now().UnixNano())
	return 0 == rand.Intn(10)
}