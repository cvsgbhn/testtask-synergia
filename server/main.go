package main

import (
	"net/http"
	"io/ioutil"
	"math/rand"
	"time"
	"fmt"
	"net"
)

func CheckUrl(url string, res chan []byte) []byte {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	res <- body
	return body
}

func Stream() {
    l, err := net.Listen("tcp", ":8080")
    if err != nil {
            fmt.Println(err)
            return
    }
    defer l.Close()

    c, err := l.Accept()
    if err != nil {
            fmt.Println(err)
            return
    }

	var c1 = make(chan []byte)
	var c2 = make(chan []byte)
	rand.Seed(time.Now().UnixNano())
	n := 0
	go func() {
		for {
			n = 1 + rand.Intn(2)
			time.Sleep(time.Second * time.Duration(n))
			CheckUrl("https://novasite.su/test1.php", c1)
		}
	}()

	go func() {
		for {
			n = 1 + rand.Intn(2)
			time.Sleep(time.Second * time.Duration(n))
			CheckUrl("https://novasite.su/test2.php", c2)
		}
	}()

	for {
		select{
		case res1 := <- c1:
			fmt.Println("from c1: ", string(res1))
			c.Write(res1)
		case res2 := <- c2:
			fmt.Println("from c2: ", string(res2))
			c.Write(res2)
		}
	}
}

func main() {
	Stream()
}