package mp

import (
    "fmt"
    "time"
)

type MP3Player struct {
    stat int
    progress int
}

func (p *MP3Player)Play(source string) {

    fmt.Println("Playing MP3 music", source)
    
    p.progress = 0

    for p.progress < 100 {
        time.Sleep(100 * time.Millisecond) // 假装正在播放
        fmt.Print(".")
        p.progress += 10
    }
    
    fmt.Println("\nFinished playing", source)
}

