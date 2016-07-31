package main

import (
	// "bufio"
	"fmt"
	"net"
	"time"
	// "strings"
)

const (
	follower = iota
	candidate
	leader
)

const (
	electionTimeout  = time.Duration(15) * time.Second
	heartbeatTimeout = time.Duration(2) * time.Second
)

const (
	heartbeat = 49		// '1' when telnet 127.0.0.1 8081
	requestVote = iota
)

type Config struct {
	state  int
	leader string
}

type Node struct {
	ip       string
	hostname string
}

func up(port string) {
	// up as follower
	fmt.Println("Launching as follower...")

	heartbeatChan := make(chan bool)
	defer close(heartbeatChan)
	heartbeatTimer := time.After(heartbeatTimeout)
	go func() {
		for {
			select {
			case <-heartbeatChan:
				fmt.Println("restart hearbeat timeout when recv msg", time.Now())
				heartbeatTimer = time.After(heartbeatTimeout)
			case <-heartbeatTimer:
				fmt.Println("hearbeat timeout", time.Now())
				heartbeatTimer = time.After(heartbeatTimeout)
			}
		}
	}()

	ln, _ := net.Listen("tcp", ":"+port)
	defer ln.Close()
	conn, _ := ln.Accept()
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		fmt.Println("msg", buf[0])
		switch buf[0] {
		case 49:
			fmt.Println("recv msg")
			heartbeatChan <- true
		default:
			// do nothing
		}

		// conn.Write(append([]byte("Message received."), buf[:]...))
	}

}

func main() {
	done := make(chan bool)

	// todo: read config to get all nodes list

	// server is up
	for i := 0; i < 1; i++ {
		go up("8081")
	}

	<-done
}
