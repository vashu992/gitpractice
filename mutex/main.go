package main

import (
	"fmt"
	"sync"
)

func main() {
	mutex := sync.Mutex{}
	accountmepaisanahihai := 0
	wg := sync.WaitGroup{}
     //wg.Add(1000)
	for i := 1; i<=100000 ; i++ {
		wg.Add(1)
        go paisabadhao(&accountmepaisanahihai , &wg , &mutex)
	}
     wg.Wait()
	 fmt.Println(accountmepaisanahihai)


}

func paisabadhao(balance *int , wg *sync.WaitGroup , tala *sync.Mutex) {
	defer wg.Done()
	// tala.Lock()
	*balance++
	// tala.Unlock()
}