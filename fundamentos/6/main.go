package main

import (
	"fmt"
)

func main() {
	slice := []int{2, 4, 6, 8, 10}
	fmt.Printf("length=%d, capacity=%d %v\n", len(slice), cap(slice), slice)

	//caso não queira ver mais items
	//quando nos colchetes tem [:numero] ou [numero:] significa que ele vai pegar os numeros após ou antes do indice que colocar
	//Se eu colocar [:2] significa que vou pegar todos os numeros antes do segundo indice que no caso é 1
	//Se eu colocar [2:] significa que vou pegar todos os numeros após o segundo indice.
	fmt.Printf("length=%d, capacity=%d %v\n", len(slice[:2]), cap(slice[:2]), slice[:2])
	fmt.Printf("length=%d, capacity=%d %v\n", len(slice[2:]), cap(slice[2:]), slice[2:])

	slice = append(slice, 12)
	fmt.Printf("length=%d, capacity=%d %v\n", len(slice), cap(slice), slice)

}
