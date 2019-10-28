package main

import ("fmt")

func main(){
	nome := "Luan"
	versao := 1.0
	
	fmt.Println(
		"Olá, Sr.", nome,
		"A versão do sistema é", versao,
	)
	fmt.Println("Escolha uma ação:")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Ver Logs")
	fmt.Println("0 - Sair")

	var input int
	fmt.Scanf("%d", &input)
	
	switch input {
		case 1:
			fmt.Println("Iniciando Monitoramento...")
		case 2:
			fmt.Println("Histórico de Logs")
		case 0:
			fmt.Println("Obrigado por utilizar nossos serviços.")
		default:
			fmt.Println("Escolha um número de 0 a 2")
	}
}