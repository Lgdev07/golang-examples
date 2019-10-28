package main

import "fmt"

func main(){
	var nome string = "Luan" 
	var idade int = 24
	var endereco string = "R Maria Lourdes" 
	var profissao string = "Programador"
	var saldo float32 = 2.300

	fmt.Println("Bom dia Sr.:", nome)
	fmt.Println("Você tem", idade, "Anos")
	fmt.Println("E mora na rua", endereco, "que legal")
	fmt.Println("Trabalha exercendo a profissão de", profissao)
	fmt.Println("Atualmente, o saldo é de:", saldo)
	fmt.Print("Atualmente, o saldo é de:", saldo)
	fmt.Print("Trabalha exercendo a profissão de", profissao)
}