package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))

	//exemplo do uso do Defer
	println("Primeira Linha")
	defer println("Segunda Linha")
	println("Terceira Linha")

}