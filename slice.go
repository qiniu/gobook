package main

import "fmt"

func main() {
	// 先定义一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// 基于数组创建一个数组切片
	var mySlice []int = myArray[:5]
	
	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	
	fmt.Println("\nElements of mySlice: ")
	
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}

	fmt.Println()
}
