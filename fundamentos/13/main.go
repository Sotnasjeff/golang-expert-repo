package main

func main() {
	var1 := 10
	var2 := 20
	println(soma(&var1, &var2))
}

func soma(a, b *int) int {
	*a = 100
	return *a + *b
}
