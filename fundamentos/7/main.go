package main

import "fmt"

func main() {

	salarios := map[string]int{"Jeff": 3000, "Trevor": 10000}
	fmt.Println(salarios["Jeff"])
	fmt.Println(salarios["Trevor"])
	delete(salarios, "Jeff")
	salarios["Michael"] = 10000

	fmt.Println(salarios["Michael"])

	//Uma forma de fazer map
	//sal := make(map[string]int)

	//Outra forma de fazer map
	//novoSal := map[string]int{}

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

}
