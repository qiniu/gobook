package main

import (
	"fmt"
)

func main() {
	var j int = 5
	
	a := func()(func()) {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j);
		}
	}()

	a()
	
	j *= 2
	
	a()
}
