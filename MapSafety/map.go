package main 

import(
	"sync"
	"fmt"
)

/* 
func main(){
	regularMap := make(map[int]interface{}) // concurrent operations when we are trying to have right to this regular map with a multiple go routines  
	// potentially having multiple go routines they're trying to write to the exact same spot in this regular map.
	syncMap := sync.Map{}

	//put
	regularMap[0] = 0 
	regularMap[1] = 1 
	regularMap[2] = 2 
	syncMap.Store(0,0)
	syncMap.Store(1,1)
	syncMap.Store(2,2)

	// get 
	regularValue, regularOk := regularMap[0]
	fmt.Println(regularValue,regularOk)
	syncValue, syncOk := syncMap.Load(0)
	fmt.Println(syncValue,syncOk)

	// delete
	regularMap[1] = nil 
	syncMap.Delete(1)

	// get and delete 

	syncValue, loaded := syncMap.LoadAndDelete(2)
	mu := sync.Mutex{}
	mu.Lock()
	regularValue = regularMap[2]
	delete(regularMap,2)
	mu.Unlock()
	fmt.Println(syncValue,loaded,regularValue)
} */

func main(){
	var syncMap sync.Map 

	syncMap.Store(1, "Value1")
	syncMap.Store(2, "Value2")

	value, exists := syncMap.Load(1)
	if exists{
		fmt.Println("Value found:",value)
	} else {
		fmt.Println("Value not found")
	}

    syncMap.Range(func(key, value interface{}) bool {
        fmt.Println("Key:", key, "Value:", value)
        return true
    })

}