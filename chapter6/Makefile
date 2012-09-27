all: echoclient echoserver hash1 https https2

echoclient: echoclient.go
	go build echoclient.go

echoserver: echoserver.go
	go build echoserver.go

hash1: hash1.go
	go build hash1.go

https: https.go
	go build https.go

https2: https2.go
	go build https2.go

clean:
	rm -f echoclient echoserver hash1 https https2
