package main

import "fmt"

func main(){
	nome := "Luan"
	versao := 1.0

	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("Escolha uma ação:")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Ver Logs")
	fmt.Println("0 - Sair")

	var comando int
	fmt.Scanf("%d", &comando)

	fmt.Println("Você escolheu a opção: ", comando)
}