package main

import "fmt"

// -------------------------------------------------------------

type IReadWriter interface {
	Read(buf *byte, cb int) int
	Write(buf *byte, cb int) int
}

// -------------------------------------------------------------

type A struct {
	a int
}

func NewA(params int) *A {
	fmt.Println("NewA:", params)
	return &A{params}
}

func (this *A) Read(buf *byte, cb int) int {
	fmt.Println("A_Read:", this.a)
	return cb
}

func (this *A) Write(buf *byte, cb int) int {
	fmt.Println("A_Write:", this.a)
	return cb
}

// -------------------------------------------------------------

type B struct {
	A
}

func NewB(params int) *B {
	fmt.Println("NewB:", params)
	return &B{A{params}}
}

func (this *B) Write(buf *byte, cb int) int {
	fmt.Println("B_Write:", this.a)
	return cb
}

func (this *B) Foo() {
	fmt.Println("B_Foo:", this.a)
}

// -------------------------------------------------------------

func main() {
	var p IReadWriter = NewB(8)
	p.Read(nil, 10)
	p.Write(nil, 10)
}

// -------------------------------------------------------------
