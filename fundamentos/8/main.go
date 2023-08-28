package main

import (
	"errors"
	"fmt"
)

func main() {

	valor, err := soma(1, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
	fmt.Println(sum(0, 1))
}

// função com retorno duplo (geralmente usado para "exceptions", pois no golang não existe exceptions)
func soma(a, b int) (int, error) {

	if a+b >= 0 {
		return a + b, errors.New("A soma é maior que 0")
	}

	return a + b, nil
}

// função normal
func sum(a, b int) int {

	if a+b >= 0 {
		return a + b
	}

	return a + b
}
