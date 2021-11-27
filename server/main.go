package main

import (
	"net/http"
	"io/ioutil"
	"time"
	"fmt"
)

func CheckUrl(url string, res chan []byte) []byte {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("for a while this is an error handler")
	}
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("for a while this is an error handler")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	res <- body
	return body
}

func main() {
	var c1 = make(chan []byte)
	var c2 = make(chan []byte)
	go func() {
		for {
			time.Sleep(time.Second * 1)
			CheckUrl("https://novasite.su/test1.php", c1)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 3)
			CheckUrl("https://novasite.su/test2.php", c2)
		}
	}()

	for {
		select{
		case res1 := <- c1:
			fmt.Println("from c1: ", string(res1))
		case res2 := <- c2:
			fmt.Println("from c2: ", string(res2))
		}
	}
}