package main

import (
	"fmt"
	"os"
)

func main(){
	
	nome:= "Luan"
	versao:= 1.0

	imprimeMenu(nome, versao)

	comandoLido := recebeValor()

	resultado(comandoLido)

}

func imprimeMenu(nome string, versao float64){
	fmt.Println(
		"Olá, Sr.", nome,
		"Essa é a versão", versao,
	)

	fmt.Println("Escolha uma ação:")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Ver Logs")
	fmt.Println("0 - Sair")
}

func recebeValor() int{
	var comando int
	fmt.Scanf("%d", &comando)
	return comando
}

func resultado(comando int){
	switch comando {
	case 1:
		fmt.Println("Iniciando monitoramento...")
	case 2:
		fmt.Println("Mostrando histórico de logs")
	case 0:
		fmt.Println("Obrigado por utilizar nossos serviços")
		os.Exit(0)
	default:
		fmt.Println("Favor escolher números de 0 a 2")
		os.Exit(-1)
}
}
// func main(){
// 	for {
// 		fmt.Println("Digite um número maior que 10:")

// 		var numero int
// 		fmt.Scanf("%d", &numero)

// 		if numero > 10 {
// 			fmt.Println("parabéns, você digitou um número maior que 10..")
// 		} else {
// 			fmt.Println("Atenção, programa sendo fechado")
// 			break
// 		}
// 	}
// }
