package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	//serializar em json
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(res))

	//criando um encoder (encoder é o cara que recebe o valor e faz o encoding em outro lugar, ele salva a informação em outro lugar)
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	//Realizando o processo de Deserialização
	jsonConta := []byte(`{"numero":2,"saldo":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonConta, &contaX)
	if err != nil {
		panic(err)
	}
	println(contaX.Saldo)

}
