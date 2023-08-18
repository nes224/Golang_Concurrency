package main

import (
	// "fmt"
	"sync"
)

func main(){
	/* 
	var beeper sync.WaitGroup
	evilNinjas := []string{"Tommy","Johnny","Bobby"}
	beeper.Add(len(evilNinjas))// added how many go routines that we want to wait 
	for _,evilNinja := range evilNinjas{
		go attack(evilNinja,&beeper)

	}
	beeper.Wait()
	fmt.Println("Mission completed") */

	var beeper sync.WaitGroup
	beeper.Add(1)
	beeper.Done()
	beeper.Wait()
}
/* 
func attack(evilNinja string, beeper *sync.WaitGroup){
	fmt.Println("Attacked evil ninja:", evilNinja)
	beeper.Done()
} */