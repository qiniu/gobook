all: closure map1 slice slice2 varg1 sorter.build

closure: closure.go
	go build closure.go

map1: map1.go
	go build map1.go

slice: slice.go
	go build slice.go

slice2: slice2.go
	go build slice2.go

varg1: varg1.go
	go build varg1.go

sorter.build:
	cd sorter; go install -v ./... 

clean:
	rm -f closure map1 slice slice2 varg1
