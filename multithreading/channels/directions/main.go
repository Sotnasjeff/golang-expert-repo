package main

// Thread 1
func main() {
	hello := make(chan string)
	go receive("Hello", hello)
	read(hello)
}

func receive(name string, hello chan<- string) { //Nessec caso se a seta estiver do lado direito do canal, significa que o canal nesse contexto irá apenas receber dados.
	hello <- name
}

func read(data <-chan string) { //Se a seta estiver no lado esquerdo, significa que o canal só está ali para esvaziar/entregar resultado
	<-data
}
