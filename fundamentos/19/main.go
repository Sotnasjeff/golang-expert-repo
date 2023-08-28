package main

func main() {

	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []int{1, 2, 3, 4, 5, 6, 10, 29, 30, 10, 30, 50}

	for i, v := range numeros {
		println(i, v)
	}

	j := 0
	for j <= 1000000 {
		println(j)
		j++
	}

}
