package main

import (
	"fmt"
	"math/rand"
)

func main(){

	numero := rand.Intn(100)
	chances := 3

	for {
		var tentativa int
		fmt.Scanf("%d", &tentativa)

		if tentativa < numero{
			chances--
			fmt.Println("O número é maior...")
			fmt.Println("Você tem mais", chances, "chances")
		} else if tentativa > numero{
			chances--
			fmt.Println("O número é menor...")
			fmt.Println("Você tem mais", chances, "chances")
		} else{
			fmt.Println("Parabéns, você ganhou, o número é:", tentativa)
			break
		}

		if chances == 0{
			fmt.Println("Suas chances acabaram, Tente novamente")
			break
		}

	}
}
