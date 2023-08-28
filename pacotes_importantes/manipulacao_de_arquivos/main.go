package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Criando arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//escrevendo com string
	//tamanho, err := f.WriteString("Hello World")
	//escrevendo com bytes
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)

	f.Close()

	//leitura
	arq, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arq))

	//leitura de pouco em pouco abrindo um arquivo
	file, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
