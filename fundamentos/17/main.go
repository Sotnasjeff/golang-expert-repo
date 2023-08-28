package main

type MyNumber int

// Você consegue criar constraints próprias
type Number interface {
	//O ~ abre uma exceção para tipos criados que são similares, como o MyNumber que é um int e o Number que pode ser Int
	~int | float64
}

func main() {
	salario := map[string]float64{"Jeff": 1000.56, "Trevor": 4000.231, "Michael": 4000.8, "Franklin": 4000.90}
	salarios := map[string]int{"Vader": 5000, "Goku": 9000, "Luke": 3000, "Thanos": 4000}
	salary := map[string]MyNumber{"Vader": 5000, "Goku": 9000, "Luke": 3000, "Thanos": 4000}

	println(Soma(salario))
	println(Soma(salarios))
	println(Soma(salary))
	println(Compara(10, 10))
}

// Uma das formas de usar genérics em Golang
func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// Comparando genérics porém o comparable não consegue identificar qual numero é maior que o outro ou menor
// Existe um pacote que
func Compara[T comparable](a T, b T) bool {
	return a == b
}

//PS: Existe o tipo any também que é para qualquer tipo, mas ele funcionaria como se fosse uma interface vazia
//Você pode fazer o mesmo exemplo de genérics para os métodos de structs
