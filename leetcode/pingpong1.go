package main

import (
    "fmt"
)

func main() {
    pingchan := make(chan string, 1)
    pongchan := make(chan string, 1)
    go ping()
    go pong()

}
