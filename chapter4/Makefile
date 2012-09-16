all: thread channel goroutine goroutine2 cgss.build

thread: thread.c
	gcc thread.c -lpthread -o thread

channel: channel.go
	go build channel.go

goroutine: goroutine.go
	go build goroutine.go

goroutine2: goroutine2.go
	go build goroutine2.go

cgss.build:
	cd cgss; go install -v ./...

clean:
	rm -f channel goroutine goroutine2 thread

