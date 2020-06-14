package main

import "fmt"

func main() {

	arrayInt := []int{5, 6, 8, 8}

	for index, value := range arrayInt {
		fmt.Println(
			"index:", index,
			"value:", value,
		)
	}

	i := 0
	for i < 5 {
		fmt.Println("I have:", i)
		i++
	}

	arrayInt = append(arrayInt, 56)

	fmt.Println(arrayInt)
}
