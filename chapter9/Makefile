all: cgo1 cgo2 primes reflect interface-all

cgo1: cgo1.go
	go build cgo1.go

cgo2: cgo2.go
	go build cgo2.go

primes: primes.go
	go build primes.go

reflect: reflect.go
	go build reflect.go

interface-all:
	cd interface; go build -o interface-1-go interface-1.go; 
	cd interface; go build -o interface-2-go interface-2.go; 
	cd interface; go build -o interface-3-go interface-3.go;  	
	cd interface; gcc interface-1.c -o interface-1-c; gcc interface-2.c -o interface-2-c; gcc interface-3.c -o interface-3-c;
clean:
	rm -f cgo1 cgo2 reflect primes interface/*-c interface/*-go
