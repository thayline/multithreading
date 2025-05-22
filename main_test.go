package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func BenchmarkRequestBrasilApi(b *testing.B) {
	req, err := http.NewRequest("GET", "http://localhost:8080/?cep=85813020", nil)
	if err != nil {
		b.Error("Erro ao criar requisição para localhost:8080")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		b.Error("Erro ao fazer requisição para localhost:8080")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		b.Error("Erro ao ler corpo da requisição para localhost:8080")
	}
	defer res.Body.Close()
	fmt.Print(body)
}

func BenchmarkRequestViaCep(b *testing.B) {
	req, err := http.NewRequest("GET", "http://localhost:8282/?cep=85813020", nil)
	if err != nil {
		b.Error("Erro ao criar requisição para localhost:8282")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		b.Error("Erro ao fazer requisição para localhost:8282")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		b.Error("Erro ao ler corpo da requisição para localhost:8282")
	}
	defer res.Body.Close()
	fmt.Print(body)
}
