package main

// Thread 1
func main() {
	forever := make(chan bool)

	// <- forever (Nesse caso, o inicio do código junto com esse pedaço que está comentado, daria um deadlock), pois ficaria esperando um evento por um tempo de forma indefinida

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()

	<-forever //Esvaziando
}
