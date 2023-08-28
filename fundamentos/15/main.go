package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
}

//Interfaces vazias podem der usado como genérics, mas a partir da versão 1.18, você não tem necessidade de usar interface vazia
//Pois o Golang ja possui um genérics
