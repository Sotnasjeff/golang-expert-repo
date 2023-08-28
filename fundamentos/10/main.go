package main

import "fmt"

func main() {

	//closure exemplo
	total := func() int {
		return sum(1, 2, 3, 3, 3, 3, 45, 6, 2, 3, 4, 1122, 12441, 5512, 551, 12, 23334, 12234, 233) * 2
	}()

	fmt.Println(total)
}

// função variádica
func sum(numeros ...int) int {
	total := 0
	for _, valor := range numeros {
		total += valor
	}

	return total
}
