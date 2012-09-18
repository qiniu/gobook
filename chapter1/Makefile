all: cprint hello paracalc reflect slice calcproj.build

cprint: cprint.go
	go build cprint.go

hello: hello.go
	go build hello.go

paracalc: paracalc.go
	go build paracalc.go

reflect: reflect.go
	go build reflect.go

slice: slice.go
	go build slice.go

calcproj.build:
	cd calcproj; go install -v ./...

clean:
	rm -f cprint hello paracalc reflect slice
