all: hello icmptest jsondemo photoweb simplehttp simplehttp2

hello: hello.go
	go build hello.go

icmptest: icmptest.go
	go build icmptest.go

jsondemo: jsondemo.go
	go build jsondemo.go

photoweb: photoweb.go
	go build photoweb.go

simplehttp: simplehttp.go
	go build simplehttp.go

simplehttp2: simplehttp2.go
	go build simplehttp2.go

clean:
	rm -f hello icmptest jsondemo photoweb simplehttp simplehttp2
