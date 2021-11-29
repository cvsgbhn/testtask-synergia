package main

import (
	"bufio"
	"time"
	"net"
	"fmt"
)

func Connect() {
    c, err := net.Dial("tcp", "127.0.0.1:8080")
    if err != nil {
            fmt.Println(err)
            return
    }

    for {
		message, _ := bufio.NewReader(c).ReadString('}')
        fmt.Println("->: ", message)
		time.Sleep(time.Second * 1)
    }
}

func main() {
	//var chn = make(chan []byte)
	Connect()
}