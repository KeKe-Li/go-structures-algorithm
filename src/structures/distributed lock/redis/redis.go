package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)



func main() {
	var counter int32
	var wg sync.WaitGroup
	//var l sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			//l.Lock()
			//counter++
			//l.Unlock()
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println(atomic.LoadInt32(&counter))
}
