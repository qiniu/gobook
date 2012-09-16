package cg

import (
    "fmt"
)

type Player struct {
    Name string
    Level int
    Exp int
    Room int

    mq chan *Message // 等待收取的消息
}

func NewPlayer() *Player {
    m := make(chan *Message, 1024)
    player := &Player{"", 0, 0, 0, m}

    go func(p *Player) {
        for {
            msg := <-p.mq
            fmt.Println(p.Name, "received message:", msg.Content)
        }
    }(player)

    return player
}
