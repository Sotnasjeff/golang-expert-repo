package main

func main() {
	jeff := Conta{
		saldo: 100,
	}

	jeff.simular(200)
	println(jeff.saldo)
}

type Conta struct {
	saldo int
}

func (c *Conta) NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}
