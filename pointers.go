package main

import "fmt"

func changeValue(pnt *string) {
	*pnt = "changed"
}

func changeValue1(pnt string) {
	pnt = "changed"
}

func main() {
	palavra := "hello"

	fmt.Println(palavra)
	changeValue(&palavra)
	fmt.Println(palavra)
}

// Ponteiros

// *variable = É o VALOR do ponteiro que a variavel está armazenando, ex:

// var strn int = 7
// var variable *int = &strn // &strn é o endereço da variável strn

// Agora a variavel variable tem o ENDEREÇO da strn como VALOR.

// fmt.Println(variable)
// -> 0xc000062150

// Podemos imprimir o valor de variable, como? Passando o * antes.
// O * é para pegarmos o VALOR de um ENDEREÇO

// Como a variable é um ENDEREÇO, agora basta buscar.

// fmt.Println(*variable)
// -> 7
