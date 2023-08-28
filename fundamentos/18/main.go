package main

import (
	"fmt"
	"main/fundamentos/18/matematica"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Car{Marca: "Chevrolet"}
	fmt.Println(carro)
	fmt.Printf("Resultado %v\n", s)
	fmt.Println(uuid.New())
}
