package main

import (
	"net/http"
	"bufio"
	"net"
	"fmt"
)

func StreamConnect(chn chan string) {
    c, err := net.Dial("tcp", "127.0.0.1:8080")
    if err != nil {
            fmt.Println(err)
            return
    }

    for {
		message, _ := bufio.NewReader(c).ReadString('}')
		chn <- message
    }
}

func SetupListening() {
	var chn = make(chan string)
	go func() {
		StreamConnect(chn)
	}()
	for {
		lastResult := <- chn
		fmt.Println(lastResult)
	}
}

func GetAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
}

func main() {
	go SetupListening()

	http.HandleFunc("/", GetAction)

    http.ListenAndServe(":8081", nil)
}