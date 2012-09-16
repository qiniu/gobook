package main

import "fmt"
import "sync"
import "runtime"

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter) //fmt.Println(z)
	lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}
