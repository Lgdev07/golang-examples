package main

import (
	"fmt"
	"./utils"
)

func main() {
	var lista = []int{5,6,8,4,6,6,8,9}

	for _, value := range lista {
		var valuesFound int
		var storedValues []int

		for _, innerValue := range lista {
			var found bool = utils.Find(innerValue, storedValues)

			if value == innerValue && !found {
				valuesFound++
			}
		}

		// fmt.Println(Find(value, storedValues))

		if valuesFound > 1 {
			storedValues = append(storedValues, value)
			fmt.Println(value)
		}
	}

}
