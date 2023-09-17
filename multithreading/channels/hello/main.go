package main

import "fmt"

func main() {
	channel := make(chan string) // Vazio

	go func() {
		channel <- "Hello World" // Cheio
	}()

	msg := <-channel // Esvaziado
	fmt.Println(msg)
}
