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

	fmt.Print("Digite o CEP: ")
	var cep string
	_, err := fmt.Scan(&cep)
	if err != nil {
		fmt.Print("Valor invalido.")
		return
	}

	c1 := make(chan []byte)
	c2 := make(chan []byte)

	go brasilApiThread(cep, c1, &wg)
	go viaCepThread(cep, c2, &wg)

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

func brasilApiThread(cep string, c1 chan<- []byte, wg *sync.WaitGroup) {

	wg.Wait()

	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)

	req, err := http.NewRequest("GET", url, nil)
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

func viaCepThread(cep string, c2 chan<- []byte, wg *sync.WaitGroup) {

	wg.Wait()

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	req, err := http.NewRequest("GET", url, nil)
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
