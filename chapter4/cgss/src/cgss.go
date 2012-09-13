package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    
    "cg"
    "ipc"
)

var centerClient *cg.CenterClient

func startCenterService() error {    

    server := ipc.NewIpcServer(&cg.CenterServer{})
    client := ipc.NewIpcClient(server)
    centerClient = &cg.CenterClient{client}

    return nil    
}

func Help(args []string) int {
    fmt.Println(`
Commands:
    login <username> <level> <exp>
    logout <username>
    send <message>
    listplayer
    quit(q)
    help(h)
`)
    return 0
}

func Quit(args []string) int {
    return 1
}

func Logout(args []string) int {
    if len(args) != 2 {
        fmt.Println("USAGE: logout <username>")
        return 0
    }
    
    centerClient.RemovePlayer(args[1])
    
    return 0
}

func Login(args []string) int {
    if len(args) != 4 {
        fmt.Println("USAGE: login <username> <level> <exp>")
        return 0
    }
    
    level, err := strconv.Atoi(args[2])
    if err != nil {
        fmt.Println("Invalid Parameter: <level> should be an integer.")
        return 0
    }
    
    exp, err := strconv.Atoi(args[3])
    if err != nil {
        fmt.Println("Invalid Parameter: <exp> should be an integer.")
        return 0
    }
    
    player := cg.NewPlayer()
    player.Name = args[1]
    player.Level = level
    player.Exp = exp
    
    err = centerClient.AddPlayer(player)
    if err != nil {
        fmt.Println("Failed adding player", err)
    }
    
    return 0
}

func ListPlayer(args []string) int {

    ps, err := centerClient.ListPlayer("")
    if err != nil {
        fmt.Println("Failed. ", err)
    } else {
        for i, v := range ps {
            fmt.Println(i + 1, ":", v)
        }
    }

    return 0
}

func Send(args []string) int {

    message := strings.Join(args[1:], " ")
    
    err := centerClient.Broadcast(message)
    if err != nil {
        fmt.Println("Failed.", err)
    }
    
    return 0
}

// 将命令和处理函数对应
func GetCommandHandlers() map[string]func(args []string) int {
    return map[string]func([]string) int {
        "help" : Help,
        "h" : Help,
        "quit" : Quit,
        "q" : Quit,
        "login" : Login,
        "logout" : Logout,
        "listplayer" : ListPlayer,
        "send" : Send,
    }
}

func main() {
    fmt.Println("Casual Game Server Solution")
        
    startCenterService()

    Help(nil)
    
    r := bufio.NewReader(os.Stdin)
    
    handlers := GetCommandHandlers()
    
    for { // 循环读取用户输入
        fmt.Print("Command> ")
        b, _, _ := r.ReadLine()
        line := string(b)
        
        tokens := strings.Split(line, " ")
        
        if handler, ok := handlers[tokens[0]]; ok {
            ret := handler(tokens)
            if ret != 0 {
                break
            }
        } else {
            fmt.Println("Unknown command:", tokens[0])
        }
    }
}
