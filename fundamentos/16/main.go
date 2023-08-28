package main

import "fmt"

func main() {
	var typeAssertion interface{} = "Hello, World"
	println(typeAssertion.(string))

	//Em casos de conversão...não é recomendável
	res, ok := typeAssertion.(int)
	fmt.Printf("O Valor de res é %v e o resultado de ok é %v\n", res, ok)
}
