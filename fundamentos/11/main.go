package main

import "fmt"

type Customer struct {
	Id       int32
	Name     string
	Age      int
	Active   bool
	Address          //Composição
	Endereco Address //Dado do tipo criado
}

// Para associar a struct, colocar o parenteses atrás com uma variavel de indetificação e o nome da struct com ou sem o ponteiro
func (c *Customer) Desativar() {
	c.Active = false
}

// O uso da Interface é implicito, uma vez que qualquer struct utilize a função de qualquer interface,
// a struct acaba implementando a interface
type Pessoa interface {
	Desativar()
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

type Address struct {
	Street string
	City   string
	State  string
}

func main() {

	jeff := Customer{
		Id:     1,
		Name:   "Jefferson",
		Age:    26,
		Active: true,
	}

	Desativacao(&jeff)

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", jeff.Name, jeff.Age, jeff.Active)

}
