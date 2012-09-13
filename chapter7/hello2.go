package main

import "fmt"

func Foo(a, b int) (ret int, err error) {
	if a > b {
		return a, nil
	} else {
		return b, nil
	}
	return 0, nil
}

func main() {
	i, _ := Foo(1, 2)
	fmt.Println("Hello, 世界", i)
}
