package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)

	c1 := make(chan []byte)
	c2 := make(chan []byte)

	go brasilApiThread(c1, &wg)
	go viaCepThread(c2, &wg)

	wg.Done()

	for {
		select {
		case res := <-c1:
			fmt.Printf("Received from brasilApi: %s\n", res)
			return

		case res := <-c2:
			fmt.Printf("Received from viaCep: %s\n", res)
			return

		case <-time.After(time.Second * 1):
			println("timeout")
		}
	}
}

func brasilApiThread(c1 chan<- []byte, wg *sync.WaitGroup) {

	wg.Wait()

	req, err := http.NewRequest("GET", "https://brasilapi.com.br/api/cep/v1/85813020", nil)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}
	c1 <- body
}

func viaCepThread(c2 chan<- []byte, wg *sync.WaitGroup) {

	wg.Wait()

	req, err := http.NewRequest("GET", "https://viacep.com.br/ws/85813020/json/", nil)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}
	c2 <- body
}
