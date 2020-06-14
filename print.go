package main

import "fmt"

func main() {
	lista := map[string]int{
		"um": 1,
		"dois": 2,
		"tres": 3,
	}

	lista2 := lista

	fmt.Printf("%p\n%p", lista, lista2)


	variable := 5
	variable2 := variable

	fmt.Printf("\n%b\n%b", variable, variable2)

}