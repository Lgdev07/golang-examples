package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	numero := r1.Intn(100)
	chances := 3

	for {
		var tentativa int
		fmt.Println("Tente adivinhar um número de 0 a 100")
		fmt.Println("%d", numero)
		fmt.Scanf("%d", &tentativa)

		if tentativa < numero {
			chances--
			fmt.Println("O número é maior...")
			fmt.Println("Você tem mais", chances, "chances")
		} else if tentativa > numero {
			chances--
			fmt.Println("O número é menor...")
			fmt.Println("Você tem mais", chances, "chances")
		} else {
			fmt.Println("Parabéns, você ganhou, o número é:", tentativa)
			break
		}

		if chances == 0 {
			fmt.Println("Suas chances acabaram, Tente novamente")
			break
		}

	}
}
