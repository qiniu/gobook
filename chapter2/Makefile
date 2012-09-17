all: closure cprint hello map1 paracalc reflect slice slice2 varg1 calcproj.build sorter.build

closure: closure.go
	go build closure.go

cprint: cprint.go
	go build cprint.go

hello: hello.go
	go build hello.go

map1: map1.go
	go build map1.go

paracalc: paracalc.go
	go build paracalc.go

reflect: reflect.go
	go build reflect.go

slice: slice.go
	go build slice.go

slice2: slice2.go
	go build slice2.go

varg1: varg1.go
	go build varg1.go

sorter.build:
	cd sorter; go install -v ./... 

calcproj.build:
	cd calcproj; go install -v ./...

clean:
	rm -f closure cprint hello map1 paracalc reflect slice slice2 varg1
