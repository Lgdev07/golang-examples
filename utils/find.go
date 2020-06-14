package utils

import "fmt"

func Find(numero int, lista []int) bool {
	for _, value := range lista {
		fmt.Println(numero, value)
		if numero == value {
			fmt.Println(numero, value)
			return true
		}
	}
	return false
}
