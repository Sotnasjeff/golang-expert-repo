package main

func main() {

	a := 10
	b := 2
	c := 3

	if a > b || c > a {
		println(a)
	} else {
		println(b)
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	case 4:
		println("f")
	default:
		println("D")
	}

}
