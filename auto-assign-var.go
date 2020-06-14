package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "teste"
	idade := 24
	saldo := 2.200

	fmt.Println(
		"Olá, seu nome é", nome,
		"Idade", idade,
		"Saldo de: ", saldo,
	)
	fmt.Println(
		"Olá, tipo de nome é", reflect.TypeOf(nome),
		"Tipo de Idade é", reflect.TypeOf(idade),
		"Tipo de Saldo é:", reflect.TypeOf(saldo),
	)
}
