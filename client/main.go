package main

import (
	"bufio"
	"time"
	"net"
	"fmt"
	"os"
)

func TCPful(chn chan []byte) {
	arguments := os.Args
    if len(arguments) == 1 {
            fmt.Println("Please provide host:port.")
            return
    }

    CONNECT := arguments[1]
    c, err := net.Dial("tcp", CONNECT)
    if err != nil {
            fmt.Println(err)
            return
    }

    for {
		message, _ := bufio.NewReader(c).ReadString('\n')
        fmt.Println("->: ", message)
		chn <- []byte(message)
		time.Sleep(time.Second * 1)
    }
}

func main() {
	var chn = make(chan []byte)
	TCPful(chn)
}