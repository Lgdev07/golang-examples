package main

import "fmt"

func MinMax(listinha []int) (int, int) {
	var max int = listinha[0]
	var min int = listinha[0]

	for _, value := range listinha {
		if value > max {
			max = value
		}

		if value < min {
			min = value
		}
	}

	return min, max
}

func main() {
	var lista = []int{5, 1, 3, 8, 4}

	min, max := MinMax(lista)

	fmt.Println(min, max)

}
