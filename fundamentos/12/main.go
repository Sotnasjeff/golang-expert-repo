package main

func main() {

	//Memória -> Endereço -> Valor

	a := 10
	var ponteiro *int = &a

	*ponteiro = 14
	println(&a)
	println(a)
	println(ponteiro)
	println(*ponteiro)

}
