package ipc

import (
    "encoding/json"
    "fmt"
)

type Request struct {
    Method string `json:"method"`
    Params string `json:"params"`
}

type Response struct {
    Code string `json:"code"`
    Body string `json:"body"`
}

type Server interface {
    Name() string
    Handle(method, params string) *Response
}

type IpcServer struct {
    Server
}

func NewIpcServer(server Server) *IpcServer {
    return &IpcServer{server}
}

func (server *IpcServer)Connect() chan string {
    session := make(chan string, 0)

    go func(c chan string) {
        for {
            request := <-c
            
            if request  == "CLOSE" { // 关闭该连接
                break
            }
            
            var req Request
            err := json.Unmarshal([]byte(request), &req)
            if err != nil {
                fmt.Println("Invalid request format:", request)
            }
            
            resp := server.Handle(req.Method, req.Params)
            
            b, err := json.Marshal(resp)
            
            c <- string(b) // 返回结果
        }
	    
	    fmt.Println("Session closed.")
	    
    }(session)
    
    fmt.Println("A new session has been created successfully.")
    
    return session
}
