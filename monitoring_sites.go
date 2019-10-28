package main

import (
	"fmt"
	"os"
	"net/http"
)

func main(){
	nome := "Luan"
	versao := 1.0

	exibeMenu(nome, versao)

	numeroEscolhido := escolheNumero()

	resultado(numeroEscolhido)
}

func exibeMenu(nome string, versao float64){
	fmt.Println(
		"Olá, Sr.", nome,
		"A versão do sistema é", versao,
	)
	fmt.Println("Escolha uma ação:")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Ver Logs")
	fmt.Println("0 - Sair")
}

func escolheNumero() int{
	var leitura int
	fmt.Scanf("%d", &leitura)
	return leitura
}

func resultado(numero int){
	switch numero{
		case 1:
			monitoraSites()		
		case 2:
			fmt.Println("Histórico de Logs")
		case 0:
			fmt.Println("Obrigado por utilizar nossos serviços")
			os.Exit(0)
		default:
			fmt.Println("Favor escolher números de 0 a 2")
			os.Exit(-1)
	}
}

func monitoraSites(){
	fmt.Println("Iniciando Monitoramento")
	site := "https://g1.globo.com/"
	teste1, teste2 := http.Get(site)
	fmt.Println(teste1)
	fmt.Println()
	fmt.Println(teste2)
}