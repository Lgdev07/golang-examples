package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	lista := []string{
		"Carro",
		"Moto",
		"JetSki",
		"Skate",
	}
	rand.Seed(time.Now().Unix())
	fmt.Println(time.Now().Unix())
	fmt.Println("Hoje vou ir ao trabalho de...", lista[rand.Intn(len(lista))])

}
