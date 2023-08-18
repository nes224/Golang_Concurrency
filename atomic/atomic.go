package main 

import(
	"fmt"
	"sync"
	"sync/atomic"
)

func main(){
	var ops uint64

	var wg  sync.WaitGroup
	for i := 0; i< 50; i++{
		wg.Add(1)

		go func(){
			defer wg.Done()
			for c :=0 ; c< 1000 ; c ++{
				// To atomically increment the counter we use AddUint64, 
				// giving it the memory address of our ops counter with the & syntax.
				atomic.AddUint64(&ops,1) 
			}
		}()
	}
	wg.Wait()

	fmt.Println("OPS:",ops)
}